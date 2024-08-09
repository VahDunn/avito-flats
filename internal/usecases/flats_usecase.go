package usecases

import (
	"avito-flats/internal/adapters/output/repositories/flats"
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"context"
)

type (
	FlatsUsecase struct {
		flats flats.Repository
	}
)

func NewFlatsUsecase(flats flats.Repository) FlatsUsecase {
	return FlatsUsecase{flats: flats}
}

func (f *FlatsUsecase) GetFlatsByHouseID(ctx context.Context, in entities.GetFlatsByHouseIDIn) ([]*entities.Flat, error) {

	var flats []*entities.Flat
	var err error

	switch in.UserType {
	case valueobjects.Client:
		flats, err = f.flats.GetFlatsByHouseID(ctx, in)
	case valueobjects.Moderator:
		flats, err = f.flats.GetFlatsByHouseIDMod(ctx, in)
	}
	if err != nil {
		return nil, err
	}

	return flats, nil

}
