package storage

import (
	"fmt"
)

const (
	// RoleBucket db
	RoleBucket string = "ROLE"
	// UserBucket db
	UserBucket string = "USER"
)

// ErrNotFound is the error returned when attempting to load a record that does
// not exist.
var ErrNotFound = fmt.Errorf("missing record")

// Bolt db
type Db interface {
	// NewID to get new id for specific bucket
	NewID(string) (string, error)
	// Save data
	Save(string, string, interface{}) error
	// Delete data
	Delete(string, string) error
	// Load data
	Load(string, string, interface{}) error
	// LoadAll data
	LoadAll(string, interface{}) error
}
