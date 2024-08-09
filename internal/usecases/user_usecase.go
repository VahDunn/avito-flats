package usecases

import (
	"avito-flats/internal/adapters/output/repositories/dummylogin"
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
)

type UserUsecase struct {
	user dummylogin.Repository
}

func NewUserUsecase(user dummylogin.Repository) UserUsecase {
	return UserUsecase{user: user}
}

func (u *UserUsecase) DummyLogin(usertype valueobjects.UserType, userid string) (entities.User, error) {

	newUser, err := u.user.DummyLogin(usertype, userid)

	if err != nil {
		return entities.User{}, err
	}

	return newUser, nil
}
