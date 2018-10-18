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

	// cemeteries section
	a.HandleFunc("/cemeteries/{id}", api.CemeteriesGetOne).Methods("GET")
	a.HandleFunc("/cemeteries/{id}", api.CemeteriesUpdate).Methods("PATCH")
	a.HandleFunc("/cemeteries", api.CemeteriesSave).Methods("POST")
	a.HandleFunc("/cemeteries", api.CemeteriesGetAll).Methods("GET")

	// contacts section
	a.HandleFunc("/contacts/{id}", api.ContactsGetOne).Methods("GET")
	a.HandleFunc("/contacts/{id}", api.ContactsUpdate).Methods("PATCH")
	a.HandleFunc("/contacts", api.ContactsSave).Methods("POST")
	a.HandleFunc("/contacts", api.ContactsGetAll).Methods("GET")

	// graves section
	a.HandleFunc("/graves/{id}", api.GravesGetOne).Methods("GET")
	a.HandleFunc("/graves/{id}", api.GravesUpdate).Methods("PATCH")
	a.HandleFunc("/graves", api.GravesSave).Methods("POST")
	a.HandleFunc("/graves", api.GravesGetAll).Methods("GET")

	return mux
}
