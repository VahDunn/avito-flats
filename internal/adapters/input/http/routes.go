package http

import (
	"avito-flats/internal/adapters/output/repositories/flats"
	"avito-flats/internal/adapters/output/repositories/flatupdate"
	"avito-flats/internal/adapters/output/repositories/house"
	"avito-flats/internal/adapters/output/repositories/newflat"
	"avito-flats/internal/usecases"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	flatsRepo := flats.InMemoryRepo{}
	flatsUsecase := usecases.NewFlatsUsecase(&flatsRepo)
	flatsHandler := NewFlatsHandler(flatsUsecase)
	router.Get("/house/{id}", flatsHandler.getFlats)

	houseRepo := house.InMemoryRepo{}
	houseUsecase := usecases.NewHouseUsecase(&houseRepo)
	houseHandler := NewHouseHandler(houseUsecase)
	router.Post("/house/create", houseHandler.createNewHouse)

	newFlatRepo := newflat.InMemoryRepo{}
	newFlatUsecase := usecases.CreateNewFlatUsecase(&newFlatRepo)
	newFlatHandler := CreateFlatHandler(newFlatUsecase)
	router.Post("/flat/create", newFlatHandler.createNewFlat)

	updateFlatRepo := flatupdate.InMemoryRepo{}
	updFlatUsecase := usecases.UpdFlatStatusUsecase(&updateFlatRepo)
	updFlatHandler := AnUpdateFlatStatusHandler(updFlatUsecase)
	router.Post("/flat/{id}/update", updFlatHandler.updateFlatStatus)

	return router
}
