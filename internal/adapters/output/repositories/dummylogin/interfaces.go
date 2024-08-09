package dummylogin

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type Repository interface {
	DummyLogin(usertype valueobjects.UserType, userid string) (entities.User, error)
}
