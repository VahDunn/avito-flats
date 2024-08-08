package subscribe

import (
	"avito-flats/internal/domain/entities"
)

type Repository interface {
	SubscribeOnHouse(ID int64, houseid entities.HouseID) (entities.HouseID, error)
}
