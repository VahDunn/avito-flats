package http

import (
	"avito-flats/internal/domain/valueobjects"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Логика обработки регистрации
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Логика обработки аутентификации
}

func (h *UserHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {

	userType := r.URL.Query().Get("user_type")

	fmt.Println(userType)
	if userType != "client" && userType != "moderator" {
		http.Error(w, "Invalid user type. Must be 'client' or 'moderator'", http.StatusBadRequest)
		return
	}

	var token valueobjects.UserType
	if userType == "moderator" {
		token = valueobjects.Moderator
	}
	if userType == "client" {
		token = valueobjects.Client
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"token": int(token),
	})
}
