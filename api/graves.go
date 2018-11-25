package api

import (
	"encoding/json"
	"github.com/catdevman/oscms-api/assemblers"
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//GravesGetOne -
func (api *API) GravesGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	result := dtos.Grave{}
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
	results := []dtos.Grave{}
	graves, err := api.graves.FindAll()

	graveAssembler := assemblers.Grave{}
	for _, g := range graves {
		result := graveAssembler.ModelToDto(&g, dtos.Grave{})
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
	graveDto := dtos.Grave{}
	err := decoder.Decode(&graveDto)

	if err != nil || graveDto.Cemetery == "" || graveDto.Location == "" {
		log.Fatal(err)
		http.Error(w, "Missing cemetery or location", http.StatusBadRequest)
		return
	}

	graveAssembler := assemblers.Grave{}
	graveModel := graveAssembler.DtoToModel(graveDto, &models.Grave{})
	grave, err := api.graves.Save(graveModel)
	if err != nil {
		DBError(w, err)
		return
	}

	result := graveAssembler.ModelToDto(grave, dtos.Grave{})

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
	jsondata := dtos.Grave{}
	_ = decoder.Decode(&jsondata)
	grave.Cemetery = bson.ObjectIdHex(jsondata.Cemetery)
	grave.Location = jsondata.Location
	err = api.graves.Update(grave)
	if err != nil {
		DBError(w, err)
		return
	}
}
