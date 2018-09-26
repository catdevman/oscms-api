package api

import (
	"github.com/catdevman/oscms-api/models"
)

// API -
type API struct {
	name       string
	cemeteries *models.CemeteryManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {
	cemeteryMgr, _ := models.NewCemeteryManager(db)
	return &API{
		name:       "OSCMS API",
		cemeteries: cemeteryMgr,
	}
}
