package api

import (
	"encoding/json"
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
	result := ContactJSON{}
	cemetery, err := api.contacts.FindContact(v["id"])
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

//ContactsGetAll -
func (api *API) ContactsGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	results := []ContactJSON{}
	contacts, err := api.contacts.FindAllContacts()
	for _, c := range contacts {
		result := AssembleContact(ContactJSON{}, &c)
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

//ContactsSave -
func (api *API) ContactsSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	decoder := json.NewDecoder(r.Body)
	jsondata := ContactJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" || jsondata.PhoneNumber == "" {
		http.Error(w, "Missing name or phoneNumber", http.StatusBadRequest)
		return
	}

	cemetery, err := api.contacts.SaveContact(jsondata.Name, jsondata.PhoneNumber)
	if err != nil {
		DBError(w, err)
		return
	}

	result := AssembleContact(ContactJSON{}, cemetery)

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
	cemetery, err := api.contacts.FindContact(v["id"])
	if err != nil {
		DBError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	jsondata := ContactJSON{}
	_ = decoder.Decode(&jsondata)
	cemetery.Name = jsondata.Name
	cemetery.PrimaryPhone = jsondata.PhoneNumber
	err = api.contacts.UpdateContact(cemetery)
	if err != nil {
		DBError(w, err)
		return
	}
}

// AssembleContact -
func AssembleContact(data ContactJSON, c *models.Contact) ContactJSON {
	data.ID = c.GetId().Hex()
	data.Name = c.Name
	data.PhoneNumber = c.PrimaryPhone
	return data
}
