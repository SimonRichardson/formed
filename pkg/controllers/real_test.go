package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		var (
			recorder = httptest.NewRecorder()
			ctrl     = New(recorder, httptest.NewRequest("GET", "/", nil))
		)

		ctrl.Get()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestPost(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		var (
			recorder = httptest.NewRecorder()
			ctrl     = New(recorder, httptest.NewRequest("POST", "/", nil))
		)

		ctrl.Post()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		var (
			recorder = httptest.NewRecorder()
			ctrl     = New(recorder, httptest.NewRequest("POST", "/bad", nil))
		)

		ctrl.NotFound()

		if expected, actual := http.StatusNotFound, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
