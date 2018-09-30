package api

import (
	"encoding/json"
	"github.com/catdevman/oscms-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

// CemeteryJSON -
type CemeteryJSON struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

//CemeteriesGetOne -
func (api *API) CemeteriesGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	result := CemeteryJSON{}
	cemetery, err := api.cemeteries.FindCemetery(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}
	result.ID = cemetery.GetId().Hex()
	result.Name = cemetery.Name
	result.PhoneNumber = cemetery.PrimaryPhone
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//CemeteriesGetAll -
func (api *API) CemeteriesGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	results := []CemeteryJSON{}
	cemeteries, err := api.cemeteries.FindAllCemeteries()
	for _, c := range cemeteries {
		result := AssembleCemetery(CemeteryJSON{}, &c)
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

//CemeteriesSave -
func (api *API) CemeteriesSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	jsondata := CemeteryJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" || jsondata.PhoneNumber == "" {
		http.Error(w, "Missing name or phoneNumber", http.StatusBadRequest)
		return
	}

	cemetery, err := api.cemeteries.SaveCemetery(jsondata.Name, jsondata.PhoneNumber)
	if err != nil {
		DBError(w, err)
		return
	}

	result := AssembleCemetery(CemeteryJSON{}, cemetery)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//CemeteriesUpdate -
func (api *API) CemeteriesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	cemetery, err := api.cemeteries.FindCemetery(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	jsondata := CemeteryJSON{}
	_ = decoder.Decode(&jsondata)
	cemetery.Name = jsondata.Name
	cemetery.PrimaryPhone = jsondata.PhoneNumber
	err = api.cemeteries.UpdateCemetery(cemetery)
	if err != nil {
		DBError(w, err)
		return
	}
}

// AssembleCemetery -
func AssembleCemetery(data CemeteryJSON, c *models.Cemetery) CemeteryJSON {
	data.ID = c.GetId().Hex()
	data.Name = c.Name
	data.PhoneNumber = c.PrimaryPhone
	return data
}
