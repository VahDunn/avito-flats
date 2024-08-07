package house

import (
	"avito-flats/internal/domain/entities"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

// PostgresRepository реализация house.Repository с использованием PostgreSQL.
type PostgresRepo struct {
	db *sql.DB
}

// Убедимся, что PostgresRepo удовлетворяет интерфейсу house.Repository.
var _ Repository = &PostgresRepo{}

var (
	NextHouseID entities.HouseID = entities.NextHouseID // Как привязать к бд?
	mutex       sync.Mutex
)

// NewPostgresRepository создает новый репозиторий для работы с базой данных.
func NewPostgresRepository(dataSourceName string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepo{db: db}, nil
}

// Close closes the database connection.
func (r *PostgresRepo) Close() error {
	if err := r.db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

func (r *PostgresRepo) CreateHouse(address string, buildyear int64, developer string) (entities.House, error) {
	// Переход к UTC и форматирование времени
	now := time.Now().UTC().Format(time.RFC3339)

	// Генерация уникального идентификатора HouseID
	mutex.Lock()
	houseId := entities.HouseID(NextHouseID)
	NextHouseID++
	mutex.Unlock()

	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return entities.House{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Вставка записи в базу данных.
	err = tx.QueryRow(`
        INSERT INTO houses (id, address, build_year, developer, creation_date, last_flat_addition_date) 
        VALUES ($1, $2, $3, $4, $5, $5) 
        RETURNING id`).Scan(&houseId)
	if err != nil {
		tx.Rollback()
		return entities.House{}, fmt.Errorf("failed to create house: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return entities.House{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Создание сущности House.
	house := entities.House{
		ID:                   houseId,
		Address:              address,
		BuildYear:            buildyear,
		Developer:            &developer,
		CreationDate:         now,
		LastFlatAdditionDate: now,
	}

	return house, nil
}
