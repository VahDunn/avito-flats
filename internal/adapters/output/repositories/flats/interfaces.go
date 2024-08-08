package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	GetFlatsByHouseID(ctx context.Context, houseID entities.HouseID) ([]*entities.Flat, error)
	GetFlatsByHouseIDMod(ctx context.Context, houseID entities.HouseID) ([]*entities.Flat, error)
}

type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
