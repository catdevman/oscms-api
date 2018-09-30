package main

import (
	"github.com/catdevman/oscms-api/api"
	"github.com/catdevman/oscms-api/models"
	"github.com/catdevman/oscms-api/routes"
	"github.com/urfave/negroni"
)

func main() {
	db, err := models.NewBongoDB("localhost", "oscms")
	if err != nil {
		panic(err)
	}
	api := api.NewAPI(db)
	r := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}
