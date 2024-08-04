package entities

import "avito-flats/internal/domain/valueobjects"

type Flat struct {
	ID               int
	HouseID          int
	Number           int
	Price            int
	RoomCount        int
	ModerationStatus valueobjects.ModerationStatus
}
