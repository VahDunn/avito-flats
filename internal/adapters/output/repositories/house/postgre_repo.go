package house

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

// CreateHouse creates a new house in the database.
func (r *PostgresRepo) CreateHouse(address string, buildYear int64, developer string) (entities.House, error) {
	now := time.Now().UTC()

	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return entities.House{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Откат транзакции в случае ошибки

	var house entities.House
	err = tx.QueryRow(`
        INSERT INTO houses (address, build_year, developer, creation_date, last_flat_addition_date) 
        VALUES ($1, $2, $3, $4, $4) 
        RETURNING id, address, build_year, developer, creation_date, last_flat_addition_date`,
		address, buildYear, developer, now).Scan(
		&house.HouseID,
		&house.Address,
		&house.BuildYear,
		&house.Developer,
		&house.CreationDate,
		&house.LastFlatAdditionDate,
	)
	if err != nil {
		return entities.House{}, fmt.Errorf("failed to create house: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		return entities.House{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return house, nil
}
