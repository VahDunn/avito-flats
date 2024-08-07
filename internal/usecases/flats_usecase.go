package usecases

import (
	"avito-flats/internal/adapters/output/repositories/flats"
	"avito-flats/internal/domain/entities"
)

type FlatsUsecase struct {
	flats flats.Repository
}

func NewFlatsUsecase(flats flats.Repository) FlatsUsecase {
	return FlatsUsecase{flats: flats}
}

func (f *FlatsUsecase) GetFlatsByHouseID(houseID entities.HouseID) ([]entities.Flat, error) {

	flats, err := f.flats.GetFlatsByHouseID(houseID)
	if err != nil {
		return nil, err
	}

	return flats, nil
}
