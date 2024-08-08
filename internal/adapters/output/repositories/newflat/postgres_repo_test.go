package newflat

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDataSourceName = "user=youruser dbname=yourtestdb sslmode=disable password=yourpassword"

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDataSourceName)
	require.NoError(t, err)

	clearTables := []string{
		"flats",
	}

	for _, table := range clearTables {
		_, err := db.Exec("DELETE FROM " + table)
		require.NoError(t, err)
	}

	return db
}

func TestPostgresRepo_CreateNewFlat(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo, err := NewPostgresRepository(testDataSourceName)
	require.NoError(t, err)
	defer repo.Close()

	houseID := entities.HouseID(1)
	number := int64(101)
	price := int64(5000000)
	roomCount := int64(3)

	newFlat, err := repo.CreateNewFlat(houseID, number, price, roomCount)
	require.NoError(t, err)

	assert.Equal(t, houseID, newFlat.HouseID)
	assert.Equal(t, number, newFlat.Number)
	assert.Equal(t, price, newFlat.Price)
	assert.Equal(t, roomCount, newFlat.RoomCount)
	assert.Equal(t, valueobjects.Created, newFlat.ModerationStatus)
}
