package query

import (
	"net/http"

	"github.com/SimonRichardson/formed/pkg/controllers"
	"github.com/SimonRichardson/formed/pkg/store"
	"github.com/SimonRichardson/formed/pkg/templates"
	"github.com/go-kit/kit/log"
)

// These are the the query API URL paths
const (
	APIPathQuery = "/"
)

// API serves the query API
type API struct {
	injector *Injector
	logger   log.Logger
}

// NewAPI creates a API with correct dependencies.
func NewAPI(injector *Injector, logger log.Logger) *API {
	return &API{
		injector: injector,
		logger:   logger,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iw := &interceptingWriter{http.StatusOK, w}
	w = iw

	// Create a new controller to handle the various routes
	ctrl := a.injector.NewController(w, r)

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

// Injector abstracts away some dependencies that are required for creating
// certain components.
type Injector struct {
	store     store.Store
	templates *templates.Templates
}

// NewInjector creates a new injector with the correct dependencies
func NewInjector(store store.Store, templates *templates.Templates) *Injector {
	return &Injector{
		store:     store,
		templates: templates,
	}
}

// NewController creates a controller from the http.ResponseWriter and the
// http.Request.
func (f *Injector) NewController(w http.ResponseWriter, r *http.Request) controllers.Controller {
	return controllers.New(f.store, f.templates, w, r)
}

type interceptingWriter struct {
	code int
	http.ResponseWriter
}

func (iw *interceptingWriter) WriteHeader(code int) {
	iw.code = code
	iw.ResponseWriter.WriteHeader(code)
}
