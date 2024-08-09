package entities

import (
	"avito-flats/internal/domain/valueobjects"
)

type User struct {
	Type   valueobjects.UserType
	UserID string
}
