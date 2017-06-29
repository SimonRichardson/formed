package models

import (
	"reflect"
	"testing"
)

func TestUserUnmarshal(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		user := &User{}
		if err := user.Unmarshal([]string{"fred", "smith"}); err != nil {
			t.Fatal(err)
		}

		if expected, actual := (User{"fred", "smith"}), *user; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("incomplete slice", func(t *testing.T) {
		user := &User{}
		err := user.Unmarshal([]string{"fred"})

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		user := &User{}
		err := user.Unmarshal(nil)

		if expected, actual := true, err != nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestUserMarshal(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		user := User{"fred", "smith"}
		slice, err := user.Marshal()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := []string{"fred", "smith"}, slice; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("empty values", func(t *testing.T) {
		user := User{"", ""}
		slice, err := user.Marshal()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := []string{"", ""}, slice; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
