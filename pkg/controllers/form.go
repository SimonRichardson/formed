package controllers

import (
	"errors"
	"net/url"

	"github.com/SimonRichardson/formed/pkg/models"
)

// UserForm creates a nice simple way to decode a form
type UserForm struct {
	FirstNames, Surnames []string
}

// DecodeFrom gets the values from a map and puts them into a more structured
// object
func (f *UserForm) DecodeFrom(values url.Values) error {
	// first decode the firstnames
	names, ok := values[formKeyFirstName]
	if !ok || len(names) == 0 {
		return errors.New("expected a series of firstnames")
	}

	f.FirstNames = names

	// secondly decode the surnames
	names, ok = values[formKeySurname]
	if !ok || len(names) == 0 {
		return errors.New("expected a series of surnames")
	}
	if len(names) != len(f.FirstNames) {
		return errors.New("expected the same number of firstnames and surnames")
	}

	f.Surnames = names

	return nil
}

// Users takes the form data and converts it into a slice of models.User. If the
// names are empty it will return an error.
func (f *UserForm) Users() ([]models.User, error) {
	users := make([]models.User, len(f.FirstNames))

	for k, v := range f.FirstNames {
		if len(v) == 0 || len(f.Surnames[k]) == 0 {
			return nil, errors.New("expected names to not be empty")
		}

		users[k] = models.User{
			FirstName: v,
			Surname:   f.Surnames[k],
		}
	}

	return users, nil
}
