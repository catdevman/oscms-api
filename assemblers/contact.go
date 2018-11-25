package assemblers

import (
	"github.com/catdevman/oscms-api/dtos"
	"github.com/catdevman/oscms-api/models"
	// "github.com/globalsign/mgo/bson"
)

// Contact -
type Contact struct {
}

// DtoToModel -
func (g *Contact) DtoToModel(contactDto dtos.Contact, contactModel *models.Contact) *models.Contact {
	if contactDto.Name != "" {
		contactModel.Name = contactDto.Name
	}

	if contactDto.PhoneNumber != "" {
		contactModel.PrimaryPhone = contactDto.PhoneNumber
	}

	return contactModel
}

// ModelToDto -
func (g *Contact) ModelToDto(contactModel *models.Contact, contactDto dtos.Contact) dtos.Contact {
	contactDto.ID = contactModel.GetId().Hex()
	contactDto.Name = contactModel.Name
	contactDto.PhoneNumber = contactModel.PrimaryPhone
	return contactDto
}
