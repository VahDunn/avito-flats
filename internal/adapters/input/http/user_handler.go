package http

import (
	"avito-flats/internal/usecases"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Логика обработки регистрации
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Логика обработки аутентификации
}
