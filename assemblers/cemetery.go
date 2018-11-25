package assemblers

import (
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	// "github.com/globalsign/mgo/bson"
)

// Cemetery -
type Cemetery struct {
}

// DtoToModel -
func (c Cemetery) DtoToModel(cemeteryDto dtos.Cemetery, cemeteryModel *models.Cemetery) *models.Cemetery {
	if cemeteryDto.Name != "" {
		cemeteryModel.Name = cemeteryDto.Name
	}

	if cemeteryDto.PhoneNumber != "" {
		cemeteryModel.PrimaryPhone = cemeteryDto.PhoneNumber
	}

	return cemeteryModel
}

// ModelToDto -
func (c Cemetery) ModelToDto(cemeteryModel *models.Cemetery, cemeteryDto dtos.Cemetery) dtos.Cemetery {
	cemeteryDto.ID = cemeteryModel.GetId().Hex()
	cemeteryDto.Name = cemeteryModel.Name
	cemeteryDto.PhoneNumber = cemeteryModel.PrimaryPhone
	return cemeteryDto
}

//
// func (c Cemetery) DtoToModel(dtos.Dto, *models.Model) *models.Model {
// 	return &models.Cemetery{}
// }
