package entities_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatCreation(t *testing.T) {
	houseID := entities.HouseID(2)
	number := int64(101)
	price := int64(1000000)
	roomCount := int64(2)
	moderationStatus := valueobjects.Created

	flat := entities.Flat{
		FlatID:           1,
		HouseID:          houseID,
		Number:           number,
		Price:            price,
		RoomCount:        roomCount,
		ModerationStatus: moderationStatus,
	}

	// Проверяем корректность создания объекта Flat
	assert.Equal(t, 1, flat.FlatID)
	assert.Equal(t, houseID, flat.HouseID)
	assert.Equal(t, number, flat.Number)
	assert.Equal(t, price, flat.Price)
	assert.Equal(t, roomCount, flat.RoomCount)
	assert.Equal(t, moderationStatus, flat.ModerationStatus)
}

func TestFlatModerationStatusChange(t *testing.T) {
	flat := entities.Flat{
		FlatID:           1,
		ModerationStatus: valueobjects.Created,
	}

	// Проверяем переход статусов модерации
	flat.ModerationStatus = valueobjects.OnModeration
	assert.Equal(t, valueobjects.OnModeration, flat.ModerationStatus, "Expected status to be OnModeration")

	flat.ModerationStatus = valueobjects.Approved
	assert.Equal(t, valueobjects.Approved, flat.ModerationStatus, "Expected status to be Approved")

	flat.ModerationStatus = valueobjects.Declined
	assert.Equal(t, valueobjects.Declined, flat.ModerationStatus, "Expected status to be Declined")
}
