package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

var _ Repository = &PostgresRepo{}

func NewPostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{db: db}
}

// GetFlatsByHouseID возвращает список квартир по ID дома
func (r *PostgresRepo) GetFlatsByHouseID(houseID entities.HouseID) ([]entities.Flat, error) {
	query := `
        SELECT id, house_id, flat_number, rooms, price, user_id
        FROM flats
        WHERE house_id = $1;
    `

	rows, err := r.db.Query(context.Background(), query, houseID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var flats []entities.Flat
	for rows.Next() {
		var flat entities.Flat
		err := rows.Scan(&flat.FlatID, &flat.HouseID, &flat.Number, &flat.RoomCount, &flat.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		flats = append(flats, flat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return flats, nil
}
