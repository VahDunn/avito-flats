package flats

import (
	"avito-flats/internal/domain/entities"
)

type Repository interface {
	GetFlatsByHouseID(houseID entities.HouseID) ([]entities.Flat, error)
}
