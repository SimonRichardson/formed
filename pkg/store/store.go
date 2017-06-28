package store

import "github.com/SimonRichardson/formed/pkg/models"

// Store is an abstraction over a underlying storage system, that allows us to
// create different implementations including mock implementation for better
// unit testing.
type Store interface {
	// Read reads all the user models from the storage, or it returns an error
	// if there issue.
	Read() ([]models.User, error)

	// Write, writes users to the underlying storage.
	Write([]models.User) error
}
