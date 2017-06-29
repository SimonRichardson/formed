package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SimonRichardson/formed/pkg/models"
	"github.com/SimonRichardson/formed/pkg/store/mock_store"
	"github.com/SimonRichardson/formed/pkg/templates"
	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate()
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("status code with no users", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, templates, recorder, httptest.NewRequest("GET", "/", nil))
		)

		store.EXPECT().
			Read().
			Return([]models.User{}, nil)

		controller.Get()

		if expected, actual := http.StatusNotFound, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("status code with some users", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, templates, recorder, httptest.NewRequest("GET", "/", nil))
		)

		store.EXPECT().
			Read().
			Return([]models.User{models.User{"Joe", "Smith"}}, nil)

		controller.Get()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("status code with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, templates, recorder, httptest.NewRequest("GET", "/", nil))
		)

		store.EXPECT().
			Read().
			Return(nil, errors.New("permissions error"))

		controller.Get()

		if expected, actual := http.StatusInternalServerError, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestPost(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate()
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, templates, recorder, httptest.NewRequest("POST", "/", nil))
		)

		controller.Post()

		if expected, actual := http.StatusOK, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate()
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store      = mock_store.NewMockStore(ctrl)
			recorder   = httptest.NewRecorder()
			controller = New(store, templates, recorder, httptest.NewRequest("POST", "/bad", nil))
		)

		controller.NotFound()

		if expected, actual := http.StatusNotFound, recorder.Code; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
