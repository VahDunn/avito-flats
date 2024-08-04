package http

import (
	"avito-flats/internal/adapters/output/repositories/flats"
	"avito-flats/internal/usecases"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	flatsRepo := flats.InMemoryRepo{}
	flatsUsecase := usecases.NewFlatsUsecase(&flatsRepo)
	flatsHandler := NewFlatsHandler(flatsUsecase)
	router.Get("/house/{id}", flatsHandler.getFlats)

	return router
}
