package main

import (
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/controller"
	"github.com/matgomes/stolen-bike-challenge/repository"
	"log"
	"net/http"
)

var Repo repository.Repository

func init() {
	Repo = repository.Repository{
		Server:   "localhost",
		Database: "police",
	}

	Repo.Connect()
}

func main() {

	handlers := controller.New(&Repo)

	router := chi.NewRouter()
	router.Get("/case", handlers.GetAll)
	router.Get("/case/{id}", handlers.GetOne)
	router.Post("/case", handlers.Create)
	router.Put("/case/{id}/close", handlers.Close)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

	defer Repo.Close()
}
