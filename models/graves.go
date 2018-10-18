package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// GravesCollection -
const GravesCollection = "graves"

// Grave -
type Grave struct {
	bongo.DocumentBase `bson:",inline"`
	Cemetery           bson.ObjectId
	Location           string
}

// GraveManager -
type GraveManager struct {
	db *DB
}

// NewGraveManager - Create a new *GraveManager that can be used for managing cemeteries.
func NewGraveManager(db *DB) (*GraveManager, error) {

	contactMgr := GraveManager{}

	contactMgr.db = db

	return &contactMgr, nil
}

// FindGrave -
func (state *GraveManager) FindGrave(id string) (*Grave, error) {
	c := &Grave{}
	err := state.db.Connection.Collection(GravesCollection).FindById(bson.ObjectIdHex(id), c)
	return c, err
}

// FindAllGraves -
func (state *GraveManager) FindAllGraves() ([]Grave, error) {
	r := []Grave{}
	c := Grave{}
	results := state.db.Connection.Collection(GravesCollection).Find(bson.M{})
	err := results.Error
	for results.Next(&c) {
		r = append(r, c)
	}
	return r, err
}

// SaveGrave -
func (state *GraveManager) SaveGrave(cemeteryID bson.ObjectId, location string) (*Grave, error) {
	g := &Grave{
		Cemetery: cemeteryID,
		Location: location,
	}
	err := state.db.Connection.Collection(GravesCollection).Save(g)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			err = vErr
		}
	}
	return g, err
}

// UpdateGrave -
func (state *GraveManager) UpdateGrave(grave *Grave) error {
	err := state.db.Collection(GravesCollection).Save(grave)
	return err
}
