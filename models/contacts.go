package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// ContactsCollection -
const ContactsCollection = "contacts"

// Contact -
type Contact struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string
	PrimaryPhone       string
}

// ContactManager -
type ContactManager struct {
	db *DB
}

// ContactRepository -
type ContactRepository interface {
	Find(id string) (*Contact, error)
	FindAll() ([]Contact, error)
	Save(contact *Contact) (*Contact, error)
	Update(contact *Contact) error
}

// NewContactManager - Create a new *ContactManager that can be used for managing cemeteries.
func NewContactManager(db *DB) (*ContactManager, error) {

	contactMgr := ContactManager{}

	contactMgr.db = db

	return &contactMgr, nil
}

// Find -
func (state *ContactManager) Find(id string) (*Contact, error) {
	c := &Contact{}
	err := state.db.Connection.Collection(ContactsCollection).FindById(bson.ObjectIdHex(id), c)
	return c, err
}

// FindAll -
func (state *ContactManager) FindAll() ([]Contact, error) {
	r := []Contact{}
	c := Contact{}
	results := state.db.Connection.Collection(ContactsCollection).Find(bson.M{})
	err := results.Error
	for results.Next(&c) {
		r = append(r, c)
	}
	return r, err
}

// Save -
func (state *ContactManager) Save(c *Contact) (*Contact, error) {
	err := state.db.Connection.Collection(ContactsCollection).Save(c)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			err = vErr
		}
	}
	return c, err
}

// Update -
func (state *ContactManager) Update(contact *Contact) error {
	err := state.db.Collection(ContactsCollection).Save(contact)
	return err
}
