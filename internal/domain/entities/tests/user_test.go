package entities_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"testing"
)

func TestUserCreation(t *testing.T) {
	// Создаем пример данных для тестирования
	userID := "124314"
	userType := valueobjects.Client // Используем тип Client

	user := entities.User{
		UserID: userID,
		Type:   userType,
	}

	// Проверяем поля пользователя
	if user.UserID != userID {
		t.Errorf("Expected user ID to be %s, got %s", userID, user.UserID)
	}

	if user.Type != userType {
		t.Errorf("Expected user type to be %d, got %d", userType, user.Type)
	}
}

// Тест для проверки типа пользователя Модератор
func TestUserCreationWithModeratorType(t *testing.T) {
	userID := "23dsgdfg"
	userType := valueobjects.Moderator // Используем тип Moderator

	user := entities.User{
		UserID: userID,
		Type:   userType,
	}

	// Проверяем поля пользователя
	if user.UserID != userID {
		t.Errorf("Expected user ID to be %s, got %s", userID, user.UserID)
	}

	if user.Type != userType {
		t.Errorf("Expected user type to be %d, got %d", userType, user.Type)
	}
}
