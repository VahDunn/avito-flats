package subscribe

import (
	"avito-flats/internal/domain/entities"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// PostgresRepo реализация house.Repository с использованием PostgreSQL.
type PostgresRepo struct {
	db *sql.DB
}

// Убедимся, что PostgresRepo удовлетворяет интерфейсу house.Repository.
var _ Repository = &PostgresRepo{}

// NewPostgresRepository создает новый репозиторий для работы с базой данных.
func NewPostgresRepository(dataSourceName string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepo{db: db}, nil
}

// Close closes the database connection.
func (r *PostgresRepo) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func (r *PostgresRepo) SubscribeOnHouse(ID int64, houseid entities.HouseID) (entities.HouseID, error) {
	now := time.Now().UTC()

	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return houseid, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Откат транзакции в случае ошибки

	var houseID int64
	err = tx.QueryRow(`
        SELECT  houses (address, build_year, developer, creation_date, last_flat_addition_date) 
        VALUES ($1, $2, $3, $4, $4) 
        RETURNING id`, address, buildYear, developer, now).Scan(&houseID)
	if err != nil {
		return houseid, fmt.Errorf("failed to find house: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		return houseid, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Создание сущности House.
	house := entities.House{
		HouseID:              entities.HouseID(houseID),
		Address:              address,
		BuildYear:            buildYear,
		Developer:            &developer,
		CreationDate:         now.Format(time.RFC3339),
		LastFlatAdditionDate: now.Format(time.RFC3339),
	}

	return house, nil
}
