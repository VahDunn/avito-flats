package house

import (
	"avito-flats/internal/domain/entities"
)

type Repository interface {
	CreateHouse(address string, buildyear int64, developer string) (entities.House, error)
}
