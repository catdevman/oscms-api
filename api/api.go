package api

import (
	"github.com/catdevman/oscms-api/models"
)

// API -
type API struct {
	name       string
	cemeteries *models.CemeteryManager
	contacts   *models.ContactManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {
	cemeteryMgr, _ := models.NewCemeteryManager(db)
	contactMgr, _ := models.NewContactManager(db)
	return &API{
		name:       "OSCMS API",
		cemeteries: cemeteryMgr,
		contacts:   contactMgr,
	}
}
