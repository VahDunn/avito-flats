package usecases

import (
	"avito-flats/internal/adapters/output/repositories"
	"avito-flats/internal/adapters/output/services"
	"avito-flats/internal/domain/valueobjects"
)

type UserUsecase struct {
	UserRepo    repositories.UserRepository
	AuthService services.AuthService
}

func (u *UserUsecase) Register(email, password string, userType valueobjects.UserType) error {
	// Логика регистрации
	return nil
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	// Логика аутентификации
	return "", nil
}
