package query

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestAPIGet(t *testing.T) {
	t.Parallel()

	t.Run("not found", func(t *testing.T) {
		var (
			api    = NewAPI(log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		res, err := request("GET", u)
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

	t.Run("not found", func(t *testing.T) {
		var (
			api    = NewAPI(log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/", server.URL)
		)
		defer server.Close()

		res, err := request("POST", u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusOK, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestAPINotFound(t *testing.T) {
	t.Parallel()

	t.Run("not found", func(t *testing.T) {
		var (
			api    = NewAPI(log.NewNopLogger())
			server = httptest.NewServer(api)

			u = fmt.Sprintf("%s/bad", server.URL)
		)
		defer server.Close()

		res, err := request("GET", u)
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := http.StatusNotFound, res.StatusCode; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func request(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	return client.Do(req)
}
