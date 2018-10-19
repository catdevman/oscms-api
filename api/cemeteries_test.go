package api

import (
	"github.com/catdevman/oscms-api/models"
	"testing"
)

// TestAssembleCemetery -
func TestAssembleCemetery(t *testing.T) {
	c := models.Cemetery{Name: "Name", PrimaryPhone: "Phone"}
	returned := CemeteryJSON{Name: "Name", PhoneNumber: "Phone"}
	r := AssembleCemetery(CemeteryJSON{}, &c)

	if r.Name != returned.Name || r.PhoneNumber != returned.PhoneNumber {
		t.Fail()
	}
}
