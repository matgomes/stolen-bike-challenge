package main

import (
	"github.com/matgomes/stolen-bike-challenge/api"
	"github.com/matgomes/stolen-bike-challenge/config"
)

func main() {
	conf := config.NewConfig("172.17.0.2", "postgres")
	a := api.NewApi(conf)
	a.SetRoutes()
	a.Start()
}
