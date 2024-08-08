package flats

import (
	"avito-flats/internal/domain/entities"
	"context"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест инициализации нового репозитория
func TestNewPostgresRepo(t *testing.T) {
	// Создаем новый pgxmock
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockPool.Close()

	repo := NewPostgresRepo(mockPool)

	assert.NotNil(t, repo)
}

// Тест метода GetFlatsByHouseID
func TestGetFlatsByHouseID(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockPool.Close()

	// Определяем запрос
	query := "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1, moderation_status = 1"
	rows := pgxmock.NewRows([]string{"id", "house_id", "flat_number", "rooms", "price"}).
		AddRow(1, 1, "101", 2, 1000000)

	mockPool.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	repo := NewPostgresRepo(mockPool)

	// Вызов метода для тестирования
	result, err := repo.GetFlatsByHouseID(context.Background(), entities.HouseID(1))
	require.NoError(t, err)

	// Проверка результатов
	expected := &entities.Flat{FlatID: 1, HouseID: 1, Number: 101, RoomCount: 2, Price: 1000000}
	assert.Len(t, result, 1)
	assert.Equal(t, expected, result[0])

	// Проверка ожиданий
	require.NoError(t, mockPool.ExpectationsWereMet())
}

// Тест метода GetFlatsByHouseIDMod
func TestGetFlatsByHouseIDMod(t *testing.T) {
	mockPool, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockPool.Close()

	// Определяем запрос
	query := "SELECT id, house_id, flat_number, rooms, price FROM flats WHERE house_id = $1"
	rows := pgxmock.NewRows([]string{"id", "house_id", "flat_number", "rooms", "price"}).
		AddRow(1, 1, "101", 2, 1000000)

	mockPool.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	repo := NewPostgresRepo(mockPool)

	// Вызов метода для тестирования
	result, err := repo.GetFlatsByHouseIDMod(context.Background(), entities.HouseID(1))
	require.NoError(t, err)

	// Проверка результатов
	expected := &entities.Flat{FlatID: 1, HouseID: 1, Number: 101, RoomCount: 2, Price: 1000000}
	assert.Len(t, result, 1)
	assert.Equal(t, expected, result[0])

	// Проверка ожиданий
	require.NoError(t, mockPool.ExpectationsWereMet())
}
