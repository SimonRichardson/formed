package query

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bytes"

	"github.com/SimonRichardson/formed/pkg/models"
	"github.com/SimonRichardson/formed/pkg/store/mock_store"
	"github.com/SimonRichardson/formed/pkg/templates"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
)

const (
	formKeyFirstName = "people[][firstname]"
	formKeySurname   = "people[][surname]"
)

func TestAPIGet(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate(false)
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("users found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store    = mock_store.NewMockStore(ctrl)
			injector = NewInjector(store, templates)
			api      = NewAPI(injector, log.NewNopLogger())
			server   = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		store.EXPECT().
			Read().
			Return([]models.User{models.User{"fred", "smith"}}, nil)

		res, err := request("GET", u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestAPIPost(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate(false)
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("valid data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store    = mock_store.NewMockStore(ctrl)
			injector = NewInjector(store, templates)
			api      = NewAPI(injector, log.NewNopLogger())
			server   = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		store.EXPECT().
			Write([]models.User{
				models.User{"fred", "bloggs"},
			}).
			Return(nil)

		formData := map[string][]string{
			formKeyFirstName: []string{"fred"},
			formKeySurname:   []string{"bloggs"},
		}

		res, err := request("POST", u, formData)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusSeeOther, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("invalid data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store    = mock_store.NewMockStore(ctrl)
			injector = NewInjector(store, templates)
			api      = NewAPI(injector, log.NewNopLogger())
			server   = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		formData := map[string][]string{
			formKeyFirstName: []string{""},
			formKeySurname:   []string{""},
		}

		res, err := request("POST", u, formData)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusBadRequest, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("no data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store    = mock_store.NewMockStore(ctrl)
			injector = NewInjector(store, templates)
			api      = NewAPI(injector, log.NewNopLogger())
			server   = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		res, err := request("POST", u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusBadRequest, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestAPINotFound(t *testing.T) {
	t.Parallel()

	fallback, err := templates.NewErrorTemplate(false)
	if err != nil {
		t.Fatal(err)
	}
	templates := templates.NewTemplates(fallback)

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			store    = mock_store.NewMockStore(ctrl)
			injector = NewInjector(store, templates)
			api      = NewAPI(injector, log.NewNopLogger())
			server   = httptest.NewServer(api)

			u = fmt.Sprintf("%s/bad", server.URL)
		)
		defer server.Close()

		res, err := request("GET", u, nil)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusNotFound, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func request(method, url string, formData map[string][]string) (*http.Response, error) {
	var body io.Reader
	if formData != nil {
		values := []string{}
		for k, v := range formData {
			for _, v := range v {
				values = append(values, fmt.Sprintf("%s=%s", k, v))
			}
		}
		body = bytes.NewBufferString(strings.Join(values, "&"))
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if formData != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return client.Do(req)
}
