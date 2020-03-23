package main

import (
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/controller"
	"github.com/matgomes/stolen-bike-challenge/repository"
	"log"
	"net/http"
)

var repo repository.Repository

func init() {
	repo = repository.Repository{
		Server:   "localhost",
		Database: "police",
	}

	repo.Connect()
}

func main() {

	handlers := controller.New(&repo)

	router := chi.NewRouter()
	router.Get("/case", handlers.GetAll)
	router.Get("/case/{id}", handlers.GetOne)
	router.Post("/case", handlers.Create)
	router.Put("/case/{id}/resolve", handlers.Resolve)

	log.Println("[ Listening on 0.0.0.0:8080 ]")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

	defer repo.Close()
}
