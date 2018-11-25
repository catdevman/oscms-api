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
	Owners             []string
	Positions          []Position
}

// GraveManager -
type GraveManager struct {
	db *DB
}

// GraveRepository -
type GraveRepository interface {
	Find(id string) (*Grave, error)
	FindAll() ([]Grave, error)
	Save(g *Grave) (*Grave, error)
	Update(grave *Grave) error
}

// NewGraveManager - Create a new GraveManager that can be used for managing cemeteries.
func NewGraveManager(db *DB) (*GraveManager, error) {

	graveMgr := GraveManager{}

	graveMgr.db = db

	return &graveMgr, nil
}

// FindGrave -
func (state *GraveManager) FindGrave(id string) (*Grave, error) {
	c := &Grave{}
	err := state.db.Connection.Collection(GravesCollection).FindById(bson.ObjectIdHex(id), c)
	return c, err
}

// FindAll -
func (state *GraveManager) FindAll() ([]Grave, error) {
	r := []Grave{}
	c := Grave{}
	results := state.db.Connection.Collection(GravesCollection).Find(bson.M{})
	err := results.Error
	for results.Next(&c) {
		r = append(r, c)
	}
	return r, err
}

// Save -
func (state *GraveManager) Save(g *Grave) (*Grave, error) {
	err := state.db.Connection.Collection(GravesCollection).Save(g)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			err = vErr
		}
	}
	return g, err
}

// Update -
func (state *GraveManager) Update(grave *Grave) error {
	err := state.db.Collection(GravesCollection).Save(grave)
	return err
}
