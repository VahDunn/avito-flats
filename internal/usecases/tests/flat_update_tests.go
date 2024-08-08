package usecases_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository это мок для интерфейса flatupdate.Repository
type MockRepository1 struct {
	mock.Mock
}

func (m *MockRepository1) UpdateFlatStatus(flatID int64, status valueobjects.ModerationStatus) (entities.Flat, error) {
	args := m.Called(flatID, status)
	return args.Get(0).(entities.Flat), args.Error(1)
}

func TestUpdateFlatStatus_Success(t *testing.T) {
	mockRepo := new(MockRepository1)
	usecase := usecases.UpdFlatStatusUsecase(mockRepo)

	flatID := int64(1)
	newStatus := valueobjects.Approved
	existingFlat := entities.Flat{
		FlatID:           flatID,
		ModerationStatus: valueobjects.OnModeration,
	}

	// Настраиваем ожидания мока
	mockRepo.On("UpdateFlatStatus", flatID, newStatus).Return(existingFlat, nil)

	updatedFlat, err := usecase.UpdateFlatStatus(flatID, newStatus)

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, flatID, updatedFlat.FlatID)
	assert.Equal(t, newStatus, updatedFlat.ModerationStatus)

	// Проверяем, что метод мока был вызван с нужными параметрами
	mockRepo.AssertExpectations(t)
}

func TestUpdateFlatStatus_Failure(t *testing.T) {
	mockRepo := new(MockRepository1)
	usecase := usecases.UpdFlatStatusUsecase(mockRepo)

	flatID := int64(1)
	newStatus := valueobjects.Approved
	expectedError := errors.New("repository update error")

	// Настраиваем ожидания мока
	mockRepo.On("UpdateFlatStatus", flatID, newStatus).Return(entities.Flat{}, expectedError)

	updatedFlat, err := usecase.UpdateFlatStatus(flatID, newStatus)

	// Проверяем результаты
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, entities.Flat{}, updatedFlat)

	// Проверяем, что метод мока был вызван с нужными параметрами
	mockRepo.AssertExpectations(t)
}
