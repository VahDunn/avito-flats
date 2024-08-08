package subscribe

import (
	"avito-flats/internal/domain/entities"
	"context"
)

type Repository interface {
	SubscribeOnHouse(ctx context.Context, HouseID entities.HouseID, email string) error
	GetSubscribers(ctx context.Context, houseID entities.HouseID) ([]string, error)
	NotifySubscribers(ctx context.Context, houseID entities.HouseID) error
}
