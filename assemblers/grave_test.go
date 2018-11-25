package assemblers

import (
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	"github.com/globalsign/mgo/bson"
	"testing"
)

//TestGraveDtoToModel -
func TestGraveDtoToModel(t *testing.T) {
	graveAssembler := Grave{}
	returned := models.Grave{
		Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"),
		Location: "Location",
		Positions: []models.Position{
			{
				Number:   1,
				Occupant: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"),
			},
		},
		Owners: []string{"1", "2"},
	}
	returned.SetId(bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"))
	graveDto := dtos.Grave{
		Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324").Hex(),
		Location: "Location",
		Positions: []dtos.Position{
			{
				Number:   1,
				Occupant: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324").Hex(),
			},
		},
		Owners: []string{"1", "2"},
	}
	r := graveAssembler.DtoToModel(graveDto, &models.Grave{})

	if r.Cemetery != returned.Cemetery || r.Location != returned.Location {
		t.Fail()
	}
}

//TestGraveModelToDto -
func TestGraveModelToDto(t *testing.T) {
	graveAssembler := Grave{}
	graveModel := models.Grave{
		Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"),
		Location: "Location",
		Positions: []models.Position{
			{
				Number:   1,
				Occupant: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"),
			},
		},
		Owners: []string{"1", "2"},
	}
	graveModel.SetId(bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324"))
	returned := dtos.Grave{
		Cemetery: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324").Hex(),
		Location: "Location",
		Positions: []dtos.Position{
			{
				Number:   1,
				Occupant: bson.ObjectIdHex("5bc7b1961ebfc51f4e7cb324").Hex(),
			},
		},
		Owners: []string{"1", "2"},
	}
	r := graveAssembler.ModelToDto(&graveModel, dtos.Grave{})

	if r.Cemetery != returned.Cemetery || r.Location != returned.Location {
		t.Fail()
	}
}
