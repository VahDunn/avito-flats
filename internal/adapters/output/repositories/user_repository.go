package repositories

import (
	"avito-flats/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(user entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}
