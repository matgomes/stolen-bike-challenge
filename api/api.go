package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/matgomes/stolen-bike-challenge/api/handler"
	"github.com/matgomes/stolen-bike-challenge/api/repository"
	"github.com/matgomes/stolen-bike-challenge/config"
	"log"
	"net/http"
)

type Api struct {
	router *httprouter.Router
	repo   *repository.Repository
}

func NewApi(conf config.Config) Api {

	db, err := repository.Connect(conf.DB)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	router := httprouter.New()

	return Api{
		router: router,
		repo:   repository.NewRepository(db),
	}

}

func (a *Api) SetRoutes() {

	a.router.GET("/case", a.handleRequest(handler.GetAllCases))
	a.router.GET("/case/:id", a.handleRequest(handler.GetOneCase))
	a.router.POST("/case", a.handleRequest(handler.CreateCase))
	a.router.POST("/case/:id/resolve", a.handleRequest(handler.ResolveCase))
}

func (a *Api) Start() {
	log.Println("[ Listening on 0.0.0.0:8080 ]")

	err := http.ListenAndServe(":8080", a.router)

	if err != nil {
		log.Fatal(err)
	}

	defer a.repo.CloseConn()
}

type RequestHandler func(r *http.Request, repo *repository.Repository, ps httprouter.Params) (code int, payload interface{})

func (a *Api) handleRequest(requestHandler RequestHandler) httprouter.Handle {

	return func(writer http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		code, payload := requestHandler(req, a.repo, ps)
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
