package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
)

type (
	PostgresRepo struct {
		db Database
	}
)

var _ Repository = &PostgresRepo{}

func NewPostgresRepo(db Database) *PostgresRepo {
	return &PostgresRepo{db: db}
}

// GetFlatsByHouseID возвращает список квартир по ID дома
func (r *PostgresRepo) GetFlatsByHouseID(ctx context.Context, in entities.GetFlatsByHouseIDIn) ([]*entities.Flat, error) {
	rows, err := r.db.Query(ctx, "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1 and moderation_status = 1", in.HouseID)
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
func (r *PostgresRepo) GetFlatsByHouseIDMod(ctx context.Context, in entities.GetFlatsByHouseIDIn) ([]*entities.Flat, error) {

	// дай мне квартиры в статусах created, approved, declined и on moderation. где я являюсь модератором
	// добавить поле в табличку квартир updated_by
	rows, err := r.db.Query(ctx, "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1 and moderation_status in (0, 2, 3) UNION SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1 and moderation_status = 1 and moderated_by = $2", in.HouseID, in.UserID)
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
