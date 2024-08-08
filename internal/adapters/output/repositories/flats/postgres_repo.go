package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
)

type PostgresRepo struct {
	db Database
}

var _ Repository = &PostgresRepo{}

func NewPostgresRepo(db Database) *PostgresRepo {
	return &PostgresRepo{db: db}
}

// GetFlatsByHouseID возвращает список квартир по ID дома
func (r *PostgresRepo) GetFlatsByHouseID(ctx context.Context, houseID entities.HouseID) ([]*entities.Flat, error) {
	rows, err := r.db.Query(ctx, "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1, moderation_status = 1", houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []*entities.Flat
	for rows.Next() {
		var flat entities.Flat
		err := rows.Scan(&flat.FlatID, &flat.HouseID, &flat.Number, &flat.RoomCount, &flat.Price)
		if err != nil {
			return nil, err
		}
		flats = append(flats, &flat)
	}
	return flats, nil
}
func (r *PostgresRepo) GetFlatsByHouseIDMod(ctx context.Context, houseID entities.HouseID) ([]*entities.Flat, error) {
	rows, err := r.db.Query(ctx, "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1", houseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []*entities.Flat
	for rows.Next() {
		var flat entities.Flat
		err := rows.Scan(&flat.FlatID, &flat.HouseID, &flat.Number, &flat.RoomCount, &flat.Price)
		if err != nil {
			return nil, err
		}
		flats = append(flats, &flat)
	}
	return flats, nil
}
