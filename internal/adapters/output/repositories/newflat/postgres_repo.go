package newflat

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

// PostgresRepository реализация house.Repository с использованием PostgreSQL.
type PostgresRepo struct {
	db *sql.DB
}

// Убедимся, что PostgresRepo удовлетворяет интерфейсу flats.Repository.
var _ Repository = &PostgresRepo{}

var (
	NextFlatID int64 = entities.NextFlatID
	mutex      sync.Mutex
)

// Создает новый репозиторий для работы с базой данных.
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

// Закрывает соединение
func (r *PostgresRepo) Close() error {
	if err := r.db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

func (r *PostgresRepo) CreateNewFlat(houseID entities.HouseID, number int64, price int64, roomcount int64) (entities.Flat, error) {
	// Переход к UTC и форматирование времени

	// Генерация FlatID
	mutex.Lock()
	flatId := NextFlatID
	entities.NextFlatID++
	mutex.Unlock()

	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return entities.Flat{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Вставка записи в базу данных.
	err = tx.QueryRow(`
        INSERT INTO flats (id, house_id, flat_number, rooms, price, user_id, moderation_status) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, 1) 
        RETURNING id`).Scan(&flatId)
	if err != nil {
		tx.Rollback()
		return entities.Flat{}, fmt.Errorf("failed to create flat: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return entities.Flat{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Создание сущности House.
	newflat := entities.Flat{
		FlatID:           flatId,
		HouseID:          houseID,
		Number:           number,
		Price:            price,
		RoomCount:        roomcount,
		ModerationStatus: valueobjects.Created,
	}

	return newflat, nil
}
