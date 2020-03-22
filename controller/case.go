package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/models"
	"github.com/matgomes/stolen-bike-challenge/repository"
	"github.com/matgomes/stolen-bike-challenge/services"
	"net/http"
)

type Controller struct {
	Service services.CaseService
}

func New(repo *repository.Repository) Controller {
	caseService := services.CaseService{Repo: repo}
	return Controller{Service: caseService}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		response, _ := json.Marshal(payload)
		w.Write(response)
	}
}

func (c *Controller) GetAll(writer http.ResponseWriter, request *http.Request) {
	all, err := c.Service.FindAllCases()

	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, nil)
		return
	}

	respondWithJson(writer, http.StatusOK, all)
}

func (c *Controller) GetOne(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	result, err := c.Service.FindCase(id)
	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, nil)
		return
	}

	respondWithJson(writer, http.StatusOK, result)
}

func (c *Controller) Create(writer http.ResponseWriter, request *http.Request) {

	var body models.Bike
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&body); err != nil {
		respondWithJson(writer, http.StatusBadRequest, nil)
		return
	}

	if err := c.Service.OpenCase(body); err != nil {
		respondWithJson(writer, http.StatusInternalServerError, nil)
		return
	}

	respondWithJson(writer, http.StatusAccepted, nil)
}

func (c *Controller) Close(writer http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")
	err := c.Service.CloseCase(id)

	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, nil)
		return
	}

	respondWithJson(writer, http.StatusOK, nil)
}
