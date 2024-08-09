package http

import (
	"avito-flats/internal/adapters/output/repositories/dummylogin"
	"avito-flats/internal/adapters/output/repositories/flats"
	"avito-flats/internal/adapters/output/repositories/flatupdate"
	"avito-flats/internal/adapters/output/repositories/house"
	"avito-flats/internal/adapters/output/repositories/newflat"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(authenticate)

	flatsRepo := flats.PostgresRepo{}
	flatsUsecase := usecases.NewFlatsUsecase(&flatsRepo)
	flatsHandler := NewFlatsHandler(flatsUsecase)
	router.Get("/house/{id}", flatsHandler.getFlats)

	houseRepo := house.PostgresRepo{}
	houseUsecase := usecases.NewHouseUsecase(&houseRepo)
	houseHandler := NewHouseHandler(houseUsecase)
	router.Post("/house/create", houseHandler.createNewHouse)

	newFlatRepo := newflat.PostgresRepo{}
	newFlatUsecase := usecases.CreateNewFlatUsecase(&newFlatRepo)
	newFlatHandler := NewFlatHandler(newFlatUsecase)
	router.Post("/flat/create", newFlatHandler.createNewFlat)

	updateFlatRepo := flatupdate.PostgresRepo{}
	updFlatUsecase := usecases.UpdFlatStatusUsecase(&updateFlatRepo)
	updFlatHandler := NewUpdateFlatStatusHandler(updFlatUsecase)
	router.Post("/flat/{id}/update", updFlatHandler.updateFlatStatus)

	userRepo := dummylogin.PostgresRepo{}
	userUsecase := usecases.NewUserUsecase(&userRepo)
	userHandler := NewUserHandler(userUsecase)
	router.Get("/dummyLogin", userHandler.DummyLogin)

	return router
}

func validateToken(token int) valueobjects.UserType {
	// This is a mock validation. In real-world application, you should validate the token properly.
	if token == 1 {
		// Return true and some user ID or info.
		return valueobjects.Moderator
	}

	return valueobjects.Client
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Bearer token validation
		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if bearerToken == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		val, err := strconv.Atoi(bearerToken)
		if err != nil {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		userType := validateToken(val)

		// Set the user ID in the context
		ctx := context.WithValue(r.Context(), "userType", userType)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
