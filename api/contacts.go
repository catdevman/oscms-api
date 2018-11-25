package api

import (
	"encoding/json"
	"github.com/catdevman/oscms-api/assemblers"
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

// ContactJSON -
type ContactJSON struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

//ContactsGetOne -
func (api *API) ContactsGetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	result := dtos.Contact{}
	contact, err := api.contacts.Find(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}
	contactAssembler := assemblers.Contact{}
	result = contactAssembler.ModelToDto(contact, result)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//ContactsGetAll -
func (api *API) ContactsGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	results := []dtos.Contact{}
	contacts, err := api.contacts.FindAll()
	if err != nil {
		DBError(w, err)
		return
	}
	contactAssembler := assemblers.Contact{}
	for _, c := range contacts {
		result := contactAssembler.ModelToDto(&c, dtos.Contact{})
		results = append(results, result)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		JSONDecodeError(w, err)
	}
}

//ContactsSave -
func (api *API) ContactsSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	jsondata := dtos.Contact{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" || jsondata.PhoneNumber == "" {
		http.Error(w, "Missing name or phoneNumber", http.StatusBadRequest)
		return
	}

	contactAssembler := assemblers.Contact{}
	contactModel := contactAssembler.DtoToModel(jsondata, &models.Contact{})
	contact, err := api.contacts.Save(contactModel)
	if err != nil {
		DBError(w, err)
		return
	}

	result := contactAssembler.ModelToDto(contact, dtos.Contact{})

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		JSONDecodeError(w, err)
		return
	}
}

//ContactsUpdate -
func (api *API) ContactsUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	v := mux.Vars(r)
	contact, err := api.contacts.Find(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	jsondata := dtos.Contact{}
	_ = decoder.Decode(&jsondata)
	contactAssembler := assemblers.Contact{}
	contact = contactAssembler.DtoToModel(jsondata, contact)
	err = api.contacts.Update(contact)
	if err != nil {
		DBError(w, err)
		return
	}
}
