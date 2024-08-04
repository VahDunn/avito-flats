package entities

import "avito-flats/internal/domain/valueobjects"

type User struct {
	ID       int
	Email    string
	Password string
	Type     valueobjects.UserType
}
