package main

import (
	"flag"
	"github.com/catdevman/oscms-api/api"
	"github.com/catdevman/oscms-api/models"
	"github.com/catdevman/oscms-api/routes"
	"github.com/urfave/negroni"
)

var port, dbHost string

func main() {
	flag.StringVar(&port, "port", "8080", "Choose a port above 1024 (default: 8080)")
	flag.StringVar(&dbHost, "databaseHost", "localhost", "Hostname/IP for MongoDB server (default: localhost)")
	flag.Parse()
	db, err := models.NewBongoDB(dbHost, "oscms")
	if err != nil {
		panic(err)
	}

	api := api.NewAPI(db)
	r := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":" + port)
}
