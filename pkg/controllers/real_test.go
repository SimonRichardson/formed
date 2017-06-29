package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SimonRichardson/formed/pkg/store/mock_store"
	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, recorder, httptest.NewRequest("GET", "/", nil))
		)

		controller.Get()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestPost(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, recorder, httptest.NewRequest("POST", "/", nil))
		)

		controller.Post()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	t.Run("status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, recorder, httptest.NewRequest("POST", "/bad", nil))
		)

		controller.NotFound()

		if expected, actual := http.StatusNotFound, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
