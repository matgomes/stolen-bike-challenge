package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/api/model"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"net/http"
	"strconv"
)

func GetAllCases(_ *http.Request, repo *repository.Repository) (int, interface{}) {

	cases, err := repo.GetAllCases()

	return handleResponse(cases, err)
}

func GetOneCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	id, _ := strconv.Atoi(chi.URLParam(request, "id"))

	result, err := repo.GetCaseByID(id)

	return handleResponse(result, err)
}

func CreateCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	var cr model.CaseRequest
	var err error

	decoder := json.NewDecoder(request.Body)

	if err = decoder.Decode(&cr); err != nil {
		return http.StatusBadRequest, nil
	}

	newCase := model.Case{CaseRequest: cr}

	officer, _ := repo.FindAvailableOfficer()

	if officer.Id.Valid {
		newCase.Officer = &officer
	}

	newCase.Id, err = repo.InsertCase(newCase, officer.Id)

	return handleResponse(newCase, err)
}

func ResolveCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	id, _ := strconv.Atoi(chi.URLParam(request, "id"))

	officerID, err := repo.ResolveCase(id)

	if err == nil && officerID.Valid {
		err = repo.AssignOpenCase(officerID)
	}

	return handleResponse(nil, err)
}

func handleResponse(result interface{}, err error) (int, interface{}) {

	if err == nil {
		return http.StatusOK, result
	}

	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil
	}

	return http.StatusInternalServerError, nil
}
