package newflat

import (
	"avito-flats/internal/domain/entities"
)

type Repository interface {
	CreateNewFlat(houseID entities.HouseID, number int64, price int64, roomcount int64) (entities.Flat, error)
}
