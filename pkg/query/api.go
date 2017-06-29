package query

import (
	"net/http"

	"github.com/SimonRichardson/formed/pkg/controllers"
	"github.com/SimonRichardson/formed/pkg/store"
	"github.com/go-kit/kit/log"
)

// These are the the query API URL paths
const (
	APIPathQuery = "/"
)

// API serves the query API
type API struct {
	facade *Facade
	logger log.Logger
}

// NewAPI creates a API with correct dependencies.
func NewAPI(facade *Facade, logger log.Logger) *API {
	return &API{
		facade: facade,
		logger: logger,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iw := &interceptingWriter{http.StatusOK, w}
	w = iw

	// Create a new controller to handle the various routes
	ctrl := a.facade.NewController(w, r)

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

// Facade abstracts away some dependencies that are required for creating
// certain components.
type Facade struct {
	store store.Store
}

// NewFacade creates a new facade with the correct dependencies
func NewFacade(store store.Store) *Facade {
	return &Facade{
		store: store,
	}
}

// NewController creates a controller from the http.ResponseWriter and the
// http.Request.
func (f *Facade) NewController(w http.ResponseWriter, r *http.Request) controllers.Controller {
	return controllers.New(f.store, w, r)
}

type interceptingWriter struct {
	code int
	http.ResponseWriter
}

func (iw *interceptingWriter) WriteHeader(code int) {
	iw.code = code
	iw.ResponseWriter.WriteHeader(code)
}
