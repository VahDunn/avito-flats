package entities

type House struct {
	ID                   HouseID
	Address              string
	BuildYear            int64
	Developer            *string
	CreationDate         string
	LastFlatAdditionDate string
}

type HouseID int64

var NextHouseID HouseID = 1
