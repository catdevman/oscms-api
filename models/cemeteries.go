package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// CemeteriesCollection -
const CemeteriesCollection = "cemeteries"

// Cemetery -
type Cemetery struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string
	PrimaryPhone       string
}

// CemeteryManager -
type CemeteryManager struct {
	db *DB
}

// CemeteryRepository -
type CemeteryRepository interface {
	Find(id string) (*Cemetery, error)
	FindAll() ([]Cemetery, error)
	Save(name, phoneNumber string) (*Cemetery, error)
	Update(cemetery *Cemetery) error
}

// NewCemeteryManager - Create a new *CemeteryManager that can be used for managing cemeteries.
func NewCemeteryManager(db *DB) (*CemeteryManager, error) {

	cemeteryMgr := CemeteryManager{}

	cemeteryMgr.db = db

	return &cemeteryMgr, nil
}

// FindCemetery -
func (state *CemeteryManager) FindCemetery(id string) (*Cemetery, error) {
	c := &Cemetery{}
	err := state.db.Connection.Collection(CemeteriesCollection).FindById(bson.ObjectIdHex(id), c)
	return c, err
}

// FindAllCemeteries -
func (state *CemeteryManager) FindAllCemeteries() ([]Cemetery, error) {
	r := []Cemetery{}
	c := Cemetery{}
	results := state.db.Connection.Collection(CemeteriesCollection).Find(bson.M{})
	err := results.Error
	for results.Next(&c) {
		r = append(r, c)
	}
	return r, err
}

// Save -
func (state *CemeteryManager) Save(c *Cemetery) (*Cemetery, error) {
	err := state.db.Connection.Collection(CemeteriesCollection).Save(c)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			err = vErr
		}
	}
	return c, err
}

// UpdateCemetery -
func (state *CemeteryManager) UpdateCemetery(cemetery *Cemetery) error {
	err := state.db.Collection(CemeteriesCollection).Save(cemetery)
	return err
}
