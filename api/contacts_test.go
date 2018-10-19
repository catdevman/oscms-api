package api

import (
	"github.com/catdevman/oscms-api/models"
	"testing"
)

//TestAssembleContact -
func TestAssembleContact(t *testing.T) {
	c := models.Contact{Name: "Name", PrimaryPhone: "Phone"}
	returned := ContactJSON{Name: "Name", PhoneNumber: "Phone"}
	r := AssembleContact(ContactJSON{}, &c)

	if r.Name != returned.Name || r.PhoneNumber != returned.PhoneNumber {
		t.Fail()
	}
}
