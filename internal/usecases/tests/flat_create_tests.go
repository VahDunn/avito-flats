package usecases_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/usecases"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNewFlatRepository - мок для репозитория newflat
type MockNewFlatRepository struct {
	mock.Mock
}

func (m *MockNewFlatRepository) CreateNewFlat(houseID entities.HouseID, number int64, price int64, roomcount int64) (entities.Flat, error) {
	args := m.Called(houseID, number, price, roomcount)
	return args.Get(0).(entities.Flat), args.Error(1)
}

func TestCreateFlatUsecase_CreateNewFlat(t *testing.T) {
	tests := []struct {
		name          string
		houseID       entities.HouseID
		number        int64
		price         int64
		roomcount     int64
		mockFlat      entities.Flat
		mockError     error
		expectedFlat  entities.Flat
		expectedError error
	}{
		{
			name:      "Successful flat creation",
			houseID:   entities.HouseID(1),
			number:    101,
			price:     100000,
			roomcount: 2,
			mockFlat: entities.Flat{
				HouseID:   entities.HouseID(1),
				Number:    101,
				Price:     100000,
				RoomCount: 2,
			},
			mockError:     nil,
			expectedFlat:  entities.Flat{HouseID: entities.HouseID(1), Number: 101, Price: 100000, RoomCount: 2},
			expectedError: nil,
		},
		{
			name:          "Error in flat creation",
			houseID:       entities.HouseID(2),
			number:        102,
			price:         150000,
			roomcount:     3,
			mockFlat:      entities.Flat{},
			mockError:     errors.New("database error"),
			expectedFlat:  entities.Flat{},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockNewFlatRepository)
			mockRepo.On("CreateNewFlat", tt.houseID, tt.number, tt.price, tt.roomcount).Return(tt.mockFlat, tt.mockError)

			usecase := usecases.CreateNewFlatUsecase(mockRepo)

			flat, err := usecase.CreateNewFlat(tt.houseID, tt.number, tt.price, tt.roomcount)

			assert.Equal(t, tt.expectedFlat, flat)
			assert.Equal(t, tt.expectedError, err)

			mockRepo.AssertExpectations(t)
		})
	}
}
