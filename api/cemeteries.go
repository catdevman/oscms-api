package api

import (
	"encoding/json"
	"github.com/catdevman/oscms-api/assemblers"
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

//CemeteriesGetOne -
func (api *API) CemeteriesGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	result := dtos.Cemetery{}
	cemetery, err := api.cemeteries.FindCemetery(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}
	cemeteryAssembler := assemblers.Cemetery{}
	result = cemeteryAssembler.ModelToDto(cemetery, result)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//CemeteriesGetAll -
func (api *API) CemeteriesGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	results := []dtos.Cemetery{}
	cemeteries, err := api.cemeteries.FindAllCemeteries()
	cemeteryAssembler := assemblers.Cemetery{}
	for _, c := range cemeteries {
		result := cemeteryAssembler.ModelToDto(&c, dtos.Cemetery{})
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
	jsondata := dtos.Cemetery{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" || jsondata.PhoneNumber == "" {
		http.Error(w, "Missing name or phoneNumber", http.StatusBadRequest)
		return
	}
	cemeteryAssembler := assemblers.Cemetery{}

	c := cemeteryAssembler.DtoToModel(jsondata, &models.Cemetery{})
	cemetery, err := api.cemeteries.Save(c)
	if err != nil {
		DBError(w, err)
		return
	}

	result := cemeteryAssembler.ModelToDto(cemetery, dtos.Cemetery{})

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
	jsondata := dtos.Cemetery{}
	_ = decoder.Decode(&jsondata)
	cemetery.Name = jsondata.Name
	cemetery.PrimaryPhone = jsondata.PhoneNumber
	err = api.cemeteries.UpdateCemetery(cemetery)
	if err != nil {
		DBError(w, err)
		return
	}
}
