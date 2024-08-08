package usecases_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository это мок для интерфейса house.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateHouse(address string, buildyear int64, developer string) (entities.House, error) {
	args := m.Called(address, buildyear, developer)
	return args.Get(0).(entities.House), args.Error(1)
}

func TestCreateHouse_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := usecases.NewHouseUsecase(mockRepo)

	address := "123 Main St"
	buildyear := int64(2000)
	developer := "Awesome Developer"
	expectedHouse := entities.House{
		HouseID:   1,
		Address:   address,
		BuildYear: buildyear,
		Developer: &developer,
	}

	mockRepo.On("CreateHouse", address, buildyear, developer).Return(expectedHouse, nil)

	result, err := usecase.CreateHouse(address, buildyear, developer)
	assert.NoError(t, err)
	assert.Equal(t, expectedHouse, result)

	mockRepo.AssertExpectations(t)
}

func TestCreateHouse_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := usecases.NewHouseUsecase(mockRepo)

	address := "123 Main St"
	buildyear := int64(2000)
	developer := "Awesome Developer"
	expectedError := errors.New("db error")

	mockRepo.On("CreateHouse", address, buildyear, developer).Return(entities.House{}, expectedError)

	result, err := usecase.CreateHouse(address, buildyear, developer)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, entities.House{}, result)

	mockRepo.AssertExpectations(t)
}
