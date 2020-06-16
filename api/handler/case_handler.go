package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/api/model"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"net/http"
	"strconv"
)

func GetAllCases(_ *http.Request, repo *repository.Repository) (int, interface{}) {

	cases, err := repo.GetAllCases()

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, cases
}

func GetOneCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	// TODO - check for error
	id, _ := strconv.Atoi(chi.URLParam(request, "id"))

	result, err := repo.GetCaseByID(id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, result
}

func CreateCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	var body model.Case
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&body); err != nil {
		return http.StatusBadRequest, nil
	}

	officer, err := repo.FindAvailableOfficer()

	if err == nil && officer.Id > 0 {
		body.Officer = officer
	}

	body.Id, err = repo.InsertCase(body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, body
}

func ResolveCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	id := chi.URLParam(request, "id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	officerID, err := repo.ResolveCase(idInt)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	go repo.UpdateUnassignedOpenCase(officerID)

	return http.StatusOK, nil
}
