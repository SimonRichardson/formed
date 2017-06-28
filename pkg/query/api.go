package query

import (
	"net/http"

	"github.com/go-kit/kit/log"
)

// These are the the query API URL paths
const (
	APIPathLoad = "/load"
	APIPathSave = "/save"
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

	// Routing table
	switch {
	default:
		http.NotFound(w, r)
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
