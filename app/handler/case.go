package handler

import (
	"encoding/json"
	"github.com/matgomes/stolen-bike-challenge/app/model"
	"github.com/matgomes/stolen-bike-challenge/app/repository"
	"net/http"
)

func GetAllCases(_ *http.Request, repo *repository.Repository) (int, interface{}) {

	cases, err := repo.GetAllCases()

	return handleResponse(cases, err)
}

func GetOneCase(request *http.Request, repo *repository.Repository) (int, interface{}) {

	id := getIDFromURLParam(request)

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

	id := getIDFromURLParam(request)

	officerID, err := repo.ResolveCase(id)

	if err == nil && officerID.Valid {
		go repo.AssignOpenCase(officerID)
	}

	return handleResponse(nil, err)
}
