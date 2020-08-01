package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/api/model"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"gopkg.in/guregu/null.v4/zero"
	"net/http"
	"strconv"
	"time"
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

	var newCase model.Case
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&newCase); err != nil {
		return http.StatusBadRequest, nil
	}

	newCase.Resolved = false
	newCase.Moment = zero.TimeFrom(time.Now())

	officer, _ := repo.FindAvailableOfficer()

	if officer.Id.Valid {
		newCase.Officer = &officer
	}

	var err error
	newCase.Id, err = repo.InsertCase(newCase, officer.Id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, newCase
}

func ResolveCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	id := chi.URLParam(request, "id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	_, err = repo.ResolveCase(idInt)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	//go repo.UpdateUnassignedOpenCase(officerID)

	return http.StatusOK, nil
}
