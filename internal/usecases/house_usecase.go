package usecases

import (
	"avito-flats/internal/adapters/output/repositories/house"
	"avito-flats/internal/domain/entities"
)

type HouseUsecase struct {
	house house.Repository
}

func NewHouseUsecase(house house.Repository) HouseUsecase {
	return HouseUsecase{house: house}
}

func (h *HouseUsecase) CreateHouse(address string, buildyear int64, developer string) (entities.House, error) {

	newHouse, err := h.house.CreateHouse(address, buildyear, developer)

	if err != nil {
		return entities.House{}, err
	}

	return newHouse, nil
}
