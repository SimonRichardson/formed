package controllers

import (
	"reflect"
	"testing"

	"github.com/SimonRichardson/formed/pkg/models"
)

func TestDecodeFrom(t *testing.T) {
	t.Parallel()

	t.Run("no data", func(t *testing.T) {
		var form UserForm
		err := form.DecodeFrom(map[string][]string{})

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("no firstnames data", func(t *testing.T) {
		var form UserForm
		err := form.DecodeFrom(map[string][]string{
			formKeyFirstName: []string{},
		})

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("no surnames data", func(t *testing.T) {
		var form UserForm
		err := form.DecodeFrom(map[string][]string{
			formKeyFirstName: []string{"fred"},
			formKeySurname:   []string{},
		})

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("misstmatch data", func(t *testing.T) {
		var form UserForm
		err := form.DecodeFrom(map[string][]string{
			formKeyFirstName: []string{"fred", "john"},
			formKeySurname:   []string{"bloggs"},
		})

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("valid data", func(t *testing.T) {
		var form UserForm
		err := form.DecodeFrom(map[string][]string{
			formKeyFirstName: []string{"fred", "john"},
			formKeySurname:   []string{"bloggs", "smith"},
		})
		if err != nil {
			t.Fatal(err)
		}

		want := UserForm{
			FirstNames: []string{"fred", "john"},
			Surnames:   []string{"bloggs", "smith"},
		}

		if expected, actual := want, form; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestUsers(t *testing.T) {
	t.Parallel()

	t.Run("invalid data", func(t *testing.T) {
		form := &UserForm{
			FirstNames: []string{""},
			Surnames:   []string{""},
		}

		_, err := form.Users()

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("valid data", func(t *testing.T) {
		form := &UserForm{
			FirstNames: []string{"fred"},
			Surnames:   []string{"bloggs"},
		}

		users, err := form.Users()
		if err != nil {
			t.Error(err)
		}

		want := []models.User{
			models.User{"fred", "bloggs"},
		}

		if expected, actual := want, users; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
