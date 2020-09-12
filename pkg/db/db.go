package storage

import (
	"fmt"
)

// ErrNotFound is the error returned when attempting to load a record that does
// not exist.
var ErrNotFound = fmt.Errorf("missing record")

// Db interface
type Db interface {
	// NewID to get new id for specific bucket
	NewID() (string, error)
	// Save data
	Save(string, interface{}) error
	// Delete data
	Delete(string) error
	// Load data
	Load(string, interface{}) error
	// LoadAll data
	LoadAll(interface{}) error
}
