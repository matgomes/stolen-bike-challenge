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

type App struct {
	mux  *chi.Mux
	repo *repository.Repository
}

type route struct {
	path    string
	handler RequestHandler
	method  string
}

func NewApp(conf config.Config) App {

	db, err := repository.Connect(conf.DB)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return App{
		mux:  chi.NewRouter(),
		repo: repository.NewRepository(db),
	}

}

func (a *App) SetRoutes() {

	routes := []route{
		{"/case", handler.GetAllCases, http.MethodGet},
		{"/case/{id:[0-9]+}", handler.GetOneCase, http.MethodGet},
		{"/case", handler.CreateCase, http.MethodPost},
		{"/case/{id:[0-9]+}/resolve", handler.ResolveCase, http.MethodPut},
	}

	configRoutes(a.mux, a.handleRequest, routes)
}

func (a *App) Start() {
	log.Println("[ Listening on 0.0.0.0:8080 ]")

	err := http.ListenAndServe(":8080", a.mux)

	if err != nil {
		log.Fatal(err)
	}

	defer a.repo.CloseConn()
}

type Middleware func(RequestHandler) http.HandlerFunc
type RequestHandler func(*http.Request, *repository.Repository) (code int, payload interface{})

func configRoutes(mux *chi.Mux, middleware Middleware, routes []route) {

	for _, r := range routes {
		mux.MethodFunc(r.method, r.path, middleware(r.handler))
	}

}

func (a *App) handleRequest(handler RequestHandler) http.HandlerFunc {

	return func(writer http.ResponseWriter, req *http.Request) {

		code, payload := handler(req, a.repo)

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
