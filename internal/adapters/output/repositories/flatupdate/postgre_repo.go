package flatupdate

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepo{db: db}, nil
}

func (r *PostgresRepo) Close() error {
	if err := r.db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

func (r *PostgresRepo) UpdateFlatStatus(flatID int64, status valueobjects.ModerationStatus) (entities.Flat, error) {
	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return entities.Flat{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        UPDATE flats 
        SET moderation_status = $2 
        WHERE id = $1`, flatID, status)
	if err != nil {
		return entities.Flat{}, fmt.Errorf("failed to update status: %w", err)
	}

	// Получение обновленной информации о квартире
	var flat entities.Flat
	err = tx.QueryRow(`
        SELECT id, house_id, flat_number, rooms, price, user_id, moderation_status
        FROM flats
        WHERE id = $1`, flatID).Scan(
		&flat.FlatID,
		&flat.HouseID,
		&flat.Number,
		&flat.RoomCount,
		&flat.Price,
		&flat.ModerationStatus,
	)
	if err != nil {
		return entities.Flat{}, fmt.Errorf("failed to fetch updated flat: %w", err)
	}

	// Подтверждение транзакции
	if err := tx.Commit(); err != nil {
		return entities.Flat{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return flat, nil
}
