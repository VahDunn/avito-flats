package usecases_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository это мок для интерфейса flats.Repository
type MockRepository2 struct {
	mock.Mock
}

func (m *MockRepository2) GetFlatsByHouseID(houseID entities.HouseID) ([]entities.Flat, error) {
	args := m.Called(houseID)
	return args.Get(0).([]entities.Flat), args.Error(1)
}

func TestGetFlatsByHouseID_Success(t *testing.T) {
	mockRepo := new(MockRepository2)
	usecase := usecases.NewFlatsUsecase(mockRepo)

	houseID := entities.HouseID(1)
	expectedFlats := []entities.Flat{
		{FlatID: 1, HouseID: houseID, RoomCount: 1, Number: 101},
		{FlatID: 2, HouseID: houseID, RoomCount: 2, Number: 102},
	}

	mockRepo.On("GetFlatsByHouseID", houseID).Return(expectedFlats, nil)

	flats, err := usecase.GetFlatsByHouseID(houseID)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlats, flats)

	mockRepo.AssertExpectations(t)
}

func TestGetFlatsByHouseID_Error(t *testing.T) {
	mockRepo := new(MockRepository2)
	usecase := usecases.NewFlatsUsecase(mockRepo)

	houseID := entities.HouseID(1)
	expectedError := errors.New("database error")

	mockRepo.On("GetFlatsByHouseID", houseID).Return(nil, expectedError)

	flats, err := usecase.GetFlatsByHouseID(houseID)
	assert.Nil(t, flats)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}
