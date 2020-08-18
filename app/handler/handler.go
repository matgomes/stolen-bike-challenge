package handler

import (
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func getIDFromURLParam(request *http.Request) int {
	id, _ := strconv.Atoi(chi.URLParam(request, "id"))
	return id
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
