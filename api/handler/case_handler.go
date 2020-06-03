package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/matgomes/stolen-bike-challenge/api/model"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"net/http"
	"strconv"
)

func GetAllCases(_ *http.Request, repo *repository.Repository, _ httprouter.Params) (int, interface{}) {

	cases, err := repo.GetAllCases()

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, cases
}

func GetOneCase(_ *http.Request, repo *repository.Repository, ps httprouter.Params) (int, interface{}) {

	id, _ := strconv.Atoi(ps.ByName("id"))

	result, err := repo.GetCaseByID(id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, result
}

func CreateCase(request *http.Request, repo *repository.Repository, _ httprouter.Params) (int, interface{}) {

	var body model.Case
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&body); err != nil {
		return http.StatusBadRequest, nil
	}

	officer, err := repo.FindAvailableOfficer()

	if err == nil {
		body.Officer = officer
	}

	body.Id, err = repo.InsertCase(body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, body
}

func ResolveCase(_ *http.Request, repo *repository.Repository, ps httprouter.Params) (int, interface{}) {

	id := ps.ByName("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	_, err = repo.ResolveCase(idInt)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	// TODO - ADD SOMETHING LIKE THIS -> go service.assignOfficerToOpenCase(officerId)

	return http.StatusOK, nil
}
