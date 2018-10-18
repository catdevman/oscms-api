package api

import (
	"github.com/catdevman/oscms-api/models"
)

// API -
type API struct {
	name       string
	cemeteries *models.CemeteryManager
	contacts   *models.ContactManager
	graves     *models.GraveManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {
	cemeteryMgr, _ := models.NewCemeteryManager(db)
	contactMgr, _ := models.NewContactManager(db)
	graveMgr, _ := models.NewGraveManager(db)
	return &API{
		name:       "OSCMS API",
		cemeteries: cemeteryMgr,
		contacts:   contactMgr,
		graves:     graveMgr,
	}
}
