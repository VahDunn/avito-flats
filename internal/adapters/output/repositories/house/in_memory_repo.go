package house

import (
	"avito-flats/internal/domain/entities"
	"fmt"
)

type InMemoryRepo struct{}

func (r *InMemoryRepo) CreateHouse(address string, buildyear int64, developer string) (entities.House, error) {
	if address == "Потешная 3" {
		return entities.House{}, fmt.Errorf("flat already exists")
	}

	newHouse := entities.House{
		ID:                   1,
		Address:              "Потешная 3",
		BuildYear:            1984,
		Developer:            nil,
		CreationDate:         "2010-11-12",
		LastFlatAdditionDate: "ssssss",
	}
	return newHouse, nil
}
