package newflat

import (
	"avito-flats/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewFlat(t *testing.T) {
	repo := &InMemoryRepo{}

	t.Run("should return an error if the flat already exists", func(t *testing.T) {
		houseID := entities.HouseID(123)
		number := int64(98)
		price := int64(10000)
		roomcount := int64(3)

		_, err := repo.CreateNewFlat(houseID, number, price, roomcount)

		assert.Error(t, err)
		assert.EqualError(t, err, "flat already exists")
	})

	t.Run("should create a new flat successfully", func(t *testing.T) {
		houseID := entities.HouseID(1)
		number := int64(44)
		price := int64(10000000000)
		roomcount := int64(2)

		flat, err := repo.CreateNewFlat(houseID, number, price, roomcount)

		assert.NoError(t, err)
		assert.Equal(t, entities.Flat{
			FlatID:           112,
			HouseID:          1,
			Number:           44,
			Price:            10000000000,
			RoomCount:        2,
			ModerationStatus: 0,
		}, flat)
	})
}
