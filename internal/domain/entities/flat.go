package entities

import "avito-flats/internal/domain/valueobjects"

type Flat struct {
	FlatID           int64
	HouseID          HouseID
	Number           int64
	Price            int64
	RoomCount        int64
	ModerationStatus valueobjects.ModerationStatus
}

var NextFlatID int64 = 1
