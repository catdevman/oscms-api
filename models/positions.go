package models

import (
	"github.com/globalsign/mgo/bson"
)

// Position -
type Position struct {
	Number   int           `json:"number"`
	Occupant bson.ObjectId `json:"occupant"`
}
