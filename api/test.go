package api

import (
	"encoding/json"
	"net/http"
)

//TestGet -
func (api *API) TestGet(w http.ResponseWriter, r *http.Request) {

	response := "ok"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
