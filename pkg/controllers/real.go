package controllers

import (
	"net/http"

	"github.com/SimonRichardson/formed/pkg/store"
)

type real struct {
	store   store.Store
	writer  http.ResponseWriter
	request *http.Request
}

// New creates a controller with the correct dependencies for the query.API
func New(store store.Store, w http.ResponseWriter, r *http.Request) Controller {
	return &real{
		store:   store,
		writer:  w,
		request: r,
	}
}

// Get defines a method for filling in the form from the store, if it finds
// nothing then it will return defaults. If an error occurs whilst
// attempting to get, then an error will be rendered.
func (r *real) Get() {
	users, err := r.store.Read()
	if err != nil {
		r.writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		r.writer.WriteHeader(http.StatusNotFound)
		return
	}
	r.writer.WriteHeader(http.StatusOK)
}

// Post consumes a form that will put the data in to the underlying store.
// If an error occurs whilst attempting to save, then an error will be
// rendered.
func (r *real) Post() {
	r.writer.WriteHeader(http.StatusOK)
}

// NotFound declares a route that doesn't exist, so an error will be
// rendered.
func (r *real) NotFound() {
	r.writer.WriteHeader(http.StatusNotFound)
}
