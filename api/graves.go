package api

import (
	"encoding/json"
	"github.com/catdevman/oscms-api/models"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"net/http"
)

// GraveJSON -
type GraveJSON struct {
	ID       string `json:"id"`
	Cemetery string `json:"cemetery"`
	Location string `json:"location"`
}

//GravesGetOne -
func (api *API) GravesGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	result := GraveJSON{}
	grave, err := api.graves.FindGrave(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}
	result.ID = grave.GetId().Hex()
	result.Cemetery = grave.Cemetery.Hex()
	result.Location = grave.Location
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//GravesGetAll -
func (api *API) GravesGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	results := []GraveJSON{}
	graves, err := api.graves.FindAllGraves()
	for _, g := range graves {
		result := AssembleGrave(GraveJSON{}, &g)
		results = append(results, result)
	}
	if err != nil {
		DBError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		JSONDecodeError(w, err)
	}
}

//GravesSave -
func (api *API) GravesSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	jsondata := GraveJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Cemetery == "" || jsondata.Location == "" {
		http.Error(w, "Missing cemetery or location", http.StatusBadRequest)
		return
	}

	grave, err := api.graves.SaveGrave(bson.ObjectIdHex(jsondata.Cemetery), jsondata.Location)
	if err != nil {
		DBError(w, err)
		return
	}

	result := AssembleGrave(GraveJSON{}, grave)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//GravesUpdate -
func (api *API) GravesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	grave, err := api.graves.FindGrave(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	jsondata := GraveJSON{}
	_ = decoder.Decode(&jsondata)
	grave.Cemetery = bson.ObjectIdHex(jsondata.Cemetery)
	grave.Location = jsondata.Location
	err = api.graves.UpdateGrave(grave)
	if err != nil {
		DBError(w, err)
		return
	}
}

// AssembleGrave -
func AssembleGrave(data GraveJSON, c *models.Grave) GraveJSON {
	data.ID = c.GetId().Hex()
	data.Cemetery = c.Cemetery.Hex()
	data.Location = c.Location
	return data
}
