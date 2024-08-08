package entities_test

import (
	"avito-flats/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseCreation(t *testing.T) {
	address := "123 Main St."
	buildYear := int64(2020)
	developer := "Some Developer"
	creationDate := "2022-01-01"
	lastFlatAdditionDate := "2022-02-01"

	house := entities.House{
		HouseID:              entities.NextHouseID,
		Address:              address,
		BuildYear:            buildYear,
		Developer:            &developer,
		CreationDate:         creationDate,
		LastFlatAdditionDate: lastFlatAdditionDate,
	}

	assert.Equal(t, entities.NextHouseID, house.HouseID)
	assert.Equal(t, address, house.Address)
	assert.Equal(t, buildYear, house.BuildYear)
	assert.Equal(t, &developer, house.Developer)
	assert.Equal(t, creationDate, house.CreationDate)
	assert.Equal(t, lastFlatAdditionDate, house.LastFlatAdditionDate)

	// Update the NextHouseID for next test
	entities.NextHouseID++
}

func TestHouseWithoutDeveloper(t *testing.T) {
	address := "456 Another St."
	buildYear := int64(2018)
	creationDate := "2021-01-01"
	lastFlatAdditionDate := "2021-02-01"

	house := entities.House{
		HouseID:              entities.NextHouseID,
		Address:              address,
		BuildYear:            buildYear,
		Developer:            nil,
		CreationDate:         creationDate,
		LastFlatAdditionDate: lastFlatAdditionDate,
	}

	assert.Equal(t, entities.NextHouseID, house.HouseID)
	assert.Equal(t, address, house.Address)
	assert.Equal(t, buildYear, house.BuildYear)
	assert.Nil(t, house.Developer)
	assert.Equal(t, creationDate, house.CreationDate)
	assert.Equal(t, lastFlatAdditionDate, house.LastFlatAdditionDate)

	// Update the NextHouseID for next test
	entities.NextHouseID++
}
