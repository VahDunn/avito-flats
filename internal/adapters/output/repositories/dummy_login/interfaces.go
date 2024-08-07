package dummylogin

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type Repository interface {
	DummyLogin(email string, usertype valueobjects.UserType) (entities.User, error)
}
