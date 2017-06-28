package models

import (
	"fmt"

	"github.com/pkg/errors"
)

// User describes a type of data that is normalized for the query API
type User struct {
	FirstName, Surname string
}

// Unmarshal converts a slice of strings to a user model
func (u *User) Unmarshal(s []string) error {
	if len(s) != 2 {
		return errors.New("expected records length of 2")
	}

	u.FirstName = s[0]
	u.Surname = s[1]

	return nil
}

// Marshal converts a user model to a slice of strings
func (u User) Marshal() ([]string, error) {
	return []string{u.FirstName, u.Surname}, nil
}

func (u User) String() string {
	return fmt.Sprintf("%s,%s", u.FirstName, u.Surname)
}
