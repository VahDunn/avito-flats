package usecases

import (
	"avito-flats/internal/adapters/output/repositories/newflat"
	"avito-flats/internal/domain/entities"
)

type CreateFlatUsecase struct {
	newflat newflat.Repository
}

func CreateNewFlatUsecase(newflat newflat.Repository) CreateFlatUsecase {
	return CreateFlatUsecase{newflat: newflat}
}

func (c *CreateFlatUsecase) CreateNewFlat(houseID entities.HouseID, number int64, price int64, roomcount int64) (entities.Flat, error) {

	newFlat, err := c.newflat.CreateNewFlat(houseID, number, price, roomcount)

	if err != nil {
		return entities.Flat{}, err
	}

	return newFlat, nil
}
