package usecases

import (
	"avito-flats/internal/adapters/output/repositories/flatupdate"
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type UpdateFlatStatusUsecase struct {
	updatedflat flatupdate.Repository
}

func UpdFlatStatusUsecase(updatedflat flatupdate.Repository) UpdateFlatStatusUsecase {
	return UpdateFlatStatusUsecase{updatedflat: updatedflat}
}

func (u *UpdateFlatStatusUsecase) UpdateFlatStatus(flatID int64, status valueobjects.ModerationStatus) (entities.Flat, error) {
	flat, err := u.updatedflat.UpdateFlatStatus(flatID, status)
	if err != nil {
		return entities.Flat{}, err
	}

	flat.ModerationStatus = status

	return flat, nil
}
