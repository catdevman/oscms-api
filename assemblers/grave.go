package assemblers

import (
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	"github.com/globalsign/mgo/bson"
)

// Grave -
type Grave struct {
}

// DtoToModel - Translate the Dto that came in as JSON into a model that can be used by the database
func (g *Grave) DtoToModel(graveDto dtos.Grave, graveModel *models.Grave) *models.Grave {
	if graveDto.Cemetery != "" {
		graveModel.Cemetery = bson.ObjectIdHex(graveDto.Cemetery)
	}

	if graveDto.Location != "" {
		graveModel.Location = graveDto.Location
	}

	if graveDto.Positions != nil {
		positions := graveModel.Positions
		for _, positionDto := range graveDto.Positions {
			position := models.Position{}
			position.Number = positionDto.Number
			position.Occupant = bson.ObjectIdHex(positionDto.Occupant)
			positions = append(positions, position)
		}
		graveModel.Positions = positions
	}

	if graveDto.Owners != nil {
		graveModel.Owners = graveDto.Owners
	}

	return graveModel
}

// ModelToDto -Translate a model to a Dto that can going across JSON
func (g *Grave) ModelToDto(graveModel *models.Grave, graveDto dtos.Grave) dtos.Grave {
	graveDto.ID = graveModel.GetId().Hex()
	graveDto.Cemetery = graveModel.Cemetery.Hex()
	graveDto.Location = graveModel.Location

	if graveModel.Positions != nil {
		var positions []dtos.Position
		for _, positionModel := range graveModel.Positions {
			position := dtos.Position{}
			position.Number = positionModel.Number
			position.Occupant = positionModel.Occupant.Hex()
			positions = append(positions, position)
		}
		graveDto.Positions = positions
	}

	graveDto.Owners = graveModel.Owners

	return graveDto
}
