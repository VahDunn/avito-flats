package main

import (
	"avito-flats/config"
	"avito-flats/internal/adapters/input/http"
	"log"
	netHttp "net/http"
)

func main() {
	cfg := config.LoadConfig()

	router := http.NewRouter()

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := netHttp.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
