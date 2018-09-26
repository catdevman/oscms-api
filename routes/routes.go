package routes

import (
	"github.com/catdevman/oscms-api/api"
	"github.com/gorilla/mux"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()
	mux.HandleFunc("/test", api.TestGet).Methods("GET")

	// api
	a := mux.PathPrefix("/api").Subrouter()

	a.HandleFunc("/cemeteries/{id}", api.CemeteriesGetOne).Methods("GET")

	a.HandleFunc("/cemeteries", api.CemeteriesSave).Methods("POST")

	a.HandleFunc("/cemeteries", api.CemeteriesGetAll).Methods("GET")

	return mux
}
