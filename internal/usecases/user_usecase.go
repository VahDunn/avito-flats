package usecases

import (
	"avito-flats/internal/adapters/output/repositories/dummy_login"
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type UserUsecase struct {
	user dummylogin.Repository
}

type Token struct {
	Token string `json:"token"`
}

func (u *UserUsecase) Register(email, password string, userType valueobjects.UserType) error {
	// Логика регистрации
	return nil
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	// Логика аутентификации
	return "", nil
}

func (u *UserUsecase) DummyLogin(email string, usertype valueobjects.UserType) (entities.User, error) {
	// Логика Dummy-аутентификации
	new_user, err := u.user.DummyLogin(email, usertype)
	if err != nil {
		return entities.User{}, err
	}
	return new_user, nil
}
