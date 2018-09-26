package models

import (
	"github.com/go-bongo/bongo"
)

// DB abstraction
type DB struct {
	*bongo.Connection
}

// NewBongoDB - mongo database
func NewBongoDB(connectionString, databaseName string) *DB {
	config := &bongo.Config{
		ConnectionString: connectionString,
		Database:         databaseName,
	}
	connection, err := bongo.Connect(config)

	if err != nil {
		panic(err)
	}

	return &DB{connection}
}
