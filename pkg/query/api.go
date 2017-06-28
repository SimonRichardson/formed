package query

import (
	"net/http"

	"github.com/SimonRichardson/formed/pkg/controllers"
	"github.com/go-kit/kit/log"
)

// These are the the query API URL paths
const (
	APIPathQuery = "/"
)

// API serves the query API
type API struct {
	logger log.Logger
}

// NewAPI creates a API with correct dependencies.
func NewAPI(logger log.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iw := &interceptingWriter{http.StatusOK, w}
	w = iw

	// Create a new controller to handle the various routes
	ctrl := controllers.New(w, r)

	// Routing table
	method, path := r.Method, r.URL.Path
	switch {
	case method == "GET" && path == APIPathQuery:
		ctrl.Get()
	case method == "POST" && path == APIPathQuery:
		ctrl.Post()
	default:
		ctrl.NotFound()
	}
}

type interceptingWriter struct {
	code int
	http.ResponseWriter
}

func (iw *interceptingWriter) WriteHeader(code int) {
	iw.code = code
	iw.ResponseWriter.WriteHeader(code)
}
