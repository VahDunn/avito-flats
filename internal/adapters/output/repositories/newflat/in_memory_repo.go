package newflat

import (
	"avito-flats/internal/domain/entities"
	"fmt"
)

type InMemoryRepo struct{}

func (r *InMemoryRepo) CreateNewFlat(houseID entities.HouseID, number int64, price int64, roomcount int64) (entities.Flat, error) {
	if houseID == 123 && number == 98 {
		return entities.Flat{}, fmt.Errorf("flat already exists")
	}

	newFlat := entities.Flat{
		FlatID:           112,
		HouseID:          1,
		Number:           44,
		Price:            10000000000,
		RoomCount:        2,
		ModerationStatus: 0,
	}
	return newFlat, nil
}
