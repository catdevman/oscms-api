package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

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

// NewContactManager - Create a new *ContactManager that can be used for managing cemeteries.
func NewContactManager(db *DB) (*ContactManager, error) {

	contactMgr := ContactManager{}

	contactMgr.db = db

	return &contactMgr, nil
}

// FindContact -
func (state *ContactManager) FindContact(id string) (*Contact, error) {
	c := &Contact{}
	err := state.db.Connection.Collection(ContactsCollection).FindById(bson.ObjectIdHex(id), c)
	return c, err
}

// FindAllContacts -
func (state *ContactManager) FindAllContacts() ([]Contact, error) {
	r := []Contact{}
	c := Contact{}
	results := state.db.Connection.Collection(ContactsCollection).Find(bson.M{})
	err := results.Error
	for results.Next(&c) {
		r = append(r, c)
	}
	return r, err
}

// SaveContact -
func (state *ContactManager) SaveContact(name, phoneNumber string) (*Contact, error) {
	c := &Contact{
		Name:         name,
		PrimaryPhone: phoneNumber,
	}
	err := state.db.Connection.Collection(ContactsCollection).Save(c)
	if err != nil {
		if vErr, ok := err.(*bongo.ValidationError); ok {
			err = vErr
		}
	}
	return c, err
}

// UpdateContact -
func (state *ContactManager) UpdateContact(contact *Contact) error {
	err := state.db.Collection(ContactsCollection).Save(contact)
	return err
}
