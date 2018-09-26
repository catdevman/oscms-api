package api

import (
	"encoding/json"
	"net/http"
)

type err struct {
	Status  string `json:"error"`
	Details string `json:"details"`
}

// JSONDecodeError -
func JSONDecodeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(err{
		Status:  "Problem Decoding JSON",
		Details: e.Error(),
	})
}

// DBError -
func DBError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(err{
		Status:  "Problem with database",
		Details: e.Error(),
	})
}
