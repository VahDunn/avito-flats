package entities

type (
	House struct {
		HouseID              HouseID
		Address              string
		BuildYear            int64
		Developer            *string
		CreationDate         string
		LastFlatAdditionDate string
	}

	HouseID int64
)

var NextHouseID HouseID = 1
