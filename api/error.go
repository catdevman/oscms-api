package api

import (
	"encoding/json"
	"net/http"
)

type err struct {
	Status  string `json:"error"`
	Details string `json:"details"`
}

// JSONDecodeErrorString -
const JSONDecodeErrorString = "Problem Decoding JSON"

// DBErrorString -
const DBErrorString = "Problem with database"

// JSONDecodeError -
func JSONDecodeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err{
		Status:  JSONDecodeErrorString,
		Details: e.Error(),
	})
}

// DBError -
func DBError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err{
		Status:  DBErrorString,
		Details: e.Error(),
	})
}
