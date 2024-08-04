package flats

import (
	"avito-flats/internal/domain/entities"
	"fmt"
)

type InMemoryRepo struct{}

func (r *InMemoryRepo) GetFlatsByHouseID(houseID entities.HouseID) ([]entities.Flat, error) {
	if houseID != 1 {
		return []entities.Flat{}, fmt.Errorf("house not found")
	}

	var flats []entities.Flat

	for i := range 5 {
		flats = append(flats, entities.Flat{
			ID:               i,
			HouseID:          1,
			Number:           i,
			Price:            1000000 * (i),
			RoomCount:        2,
			ModerationStatus: 0,
		})
	}

	return flats, nil
}
