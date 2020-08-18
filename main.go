package main

import (
	"github.com/matgomes/stolen-bike-challenge/app"
	"github.com/matgomes/stolen-bike-challenge/config"
)

func main() {
	conf := config.NewConfig("172.17.0.2", "postgres")
	a := app.NewApp(conf)
	a.SetRoutes()
	a.Start()
}
