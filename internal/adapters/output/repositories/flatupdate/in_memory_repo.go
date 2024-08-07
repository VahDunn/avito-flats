package flatupdate

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"fmt"
)

type InMemoryRepo struct{}

func (r *InMemoryRepo) UpdateFlatStatus(flatID int64, status valueobjects.ModerationStatus) (entities.Flat, error) {
	if flatID != 1 {
		return entities.Flat{}, fmt.Errorf("house not found")
	}

	var flat entities.Flat
	{
		flat = entities.Flat{
			FlatID:           1,
			HouseID:          1,
			Number:           5,
			Price:            500000000,
			RoomCount:        23,
			ModerationStatus: 1,
		}

	}

	return flat, nil
}
