//go:generate mockgen -destination=./mocks/interfaces.go -package=${GOPACKAGE} -source=interfaces.go
package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
	"github.com/jackc/pgx/v4"
)

type (
	Repository interface {
		GetFlatsByHouseID(ctx context.Context, in entities.GetFlatsByHouseIDIn) ([]*entities.Flat, error)
		GetFlatsByHouseIDMod(ctx context.Context, in entities.GetFlatsByHouseIDIn) ([]*entities.Flat, error)
	}

	Database interface {
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	}
)
