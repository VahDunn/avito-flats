package usecases

import (
	"avito-flats/internal/adapters/output/repositories/subscribe"
	"avito-flats/internal/domain/entities"
)

type SubscribeUsecase struct {
	subscribe subscribe.Repository
}

func SubscribeUsecase(house subscribe.Repository) SubscribeUsecase {
	return SubscribeUsecase{subscribe: subscribe}
}

func (h *SubscribeUscase) SubscribeOnHouse(ID int64, houseid entities.HouseID) (entities.HouseID, error) {

	newSubscribe, err := h.subscribe.SubscribeOnHouse(ID, houseid)

	if err != nil {
		return string, err
	}

	return newSubscribe, nil
}
