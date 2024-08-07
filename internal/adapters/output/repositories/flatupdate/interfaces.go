package flatupdate

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type Repository interface {
	UpdateFlatStatus(flatID int64, status valueobjects.ModerationStatus) (entities.Flat, error)
}
