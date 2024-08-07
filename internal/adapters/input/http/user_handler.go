package http

import (
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func (h *UserHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	// Логика упрощенной аутентификации и авторизации
	userType := r.URL.Query().Get("user_type")
	if userType == "" {
		http.Error(w, "user_type is required", http.StatusBadRequest)
		return
	}

	// Генерация упрощенного токена (используя текущее время)
	tokenStr := fmt.Sprintf("dummy_token_%s_%d", userType, time.Now().UnixNano())

	token := valueobjects.Token{
		Token: tokenStr,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
