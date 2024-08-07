package http

import (
	"avito-flats/internal/usecases"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Логика обработки регистрации
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Логика обработки аутентификации
}

func (h *UserHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserType string `json:"user_type"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserType != "client" && req.UserType != "moderator" {
		http.Error(w, "Invalid user type. Must be 'client' or 'moderator'", http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserType: req.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
