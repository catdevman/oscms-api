package api

import (
	"github.com/catdevman/oscms-api/models"
	"github.com/globalsign/mgo/bson"
	"testing"
)

//TestAssembleGrave -
func TestAssembleGrave(t *testing.T) {
	c := models.Grave{Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"), Location: "Location"}
	c.SetId(bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"))
	returned := GraveJSON{Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324").Hex(), Location: "Location"}
	r := AssembleGrave(GraveJSON{}, &c)

	if r.Cemetery != returned.Cemetery || r.Location != returned.Location {
		t.Fail()
	}
}
