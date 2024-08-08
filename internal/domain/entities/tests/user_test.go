package entities_test

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreation(t *testing.T) {
	id := int64(1)
	email := "testuser@example.com"
	password := "securepassword"
	userType := valueobjects.Client

	user := entities.User{
		ID:       id,
		Email:    email,
		Password: password,
		Type:     userType,
	}

	// Проверяем корректность создания объекта User
	assert.Equal(t, id, user.ID)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, userType, user.Type)
}

func TestDifferentUserTypes(t *testing.T) {
	client := entities.User{
		ID:       2,
		Email:    "client@example.com",
		Password: "clientpassword",
		Type:     valueobjects.Client,
	}

	moderator := entities.User{
		ID:       3,
		Email:    "moderator@example.com",
		Password: "moderatorpassword",
		Type:     valueobjects.Moderator,
	}

	// Проверка клиента
	assert.Equal(t, int64(2), client.ID)
	assert.Equal(t, "client@example.com", client.Email)
	assert.Equal(t, "clientpassword", client.Password)
	assert.Equal(t, valueobjects.Client, client.Type)

	// Проверка модератора
	assert.Equal(t, int64(3), moderator.ID)
	assert.Equal(t, "moderator@example.com", moderator.Email)
	assert.Equal(t, "moderatorpassword", moderator.Password)
	assert.Equal(t, valueobjects.Moderator, moderator.Type)
}
