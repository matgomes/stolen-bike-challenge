package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/matgomes/stolen-bike-challenge/api/handler"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"github.com/matgomes/stolen-bike-challenge/config"
	"log"
	"net/http"
)

type Api struct {
	mux  *chi.Mux
	repo *repository.Repository
}

func NewApi(conf config.Config) Api {

	db, err := repository.Connect(conf.DB)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return Api{
		mux:  chi.NewRouter(),
		repo: repository.NewRepository(db),
	}

}

func (a *Api) SetRoutes() {

	a.mux.Get("/case", a.handleRequest(handler.GetAllCases))
	a.mux.Get("/case/{id}", a.handleRequest(handler.GetOneCase))
	a.mux.Post("/case", a.handleRequest(handler.CreateCase))
	a.mux.Post("/case/{id}/resolve", a.handleRequest(handler.ResolveCase))
}

func (a *Api) Start() {
	log.Println("[ Listening on 0.0.0.0:8080 ]")

	err := http.ListenAndServe(":8080", a.mux)

	if err != nil {
		log.Fatal(err)
	}

	defer a.repo.CloseConn()
}

type RequestHandler func(r *http.Request, repo *repository.Repository) (code int, payload interface{})

func (a *Api) handleRequest(requestHandler RequestHandler) http.HandlerFunc {

	return func(writer http.ResponseWriter, req *http.Request) {
		code, payload := requestHandler(req, a.repo)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(code)

		if payload == nil {
			return
		}

		response, err := json.Marshal(payload)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		writer.Write(response)
	}

}
