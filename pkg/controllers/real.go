package controllers

import (
	"errors"
	"net/http"

	"github.com/SimonRichardson/formed/pkg/store"
	"github.com/SimonRichardson/formed/pkg/templates"
)

type real struct {
	store     store.Store
	templates *templates.Templates
	writer    http.ResponseWriter
	request   *http.Request
}

// New creates a controller with the correct dependencies for the query.API
func New(s store.Store, t *templates.Templates, w http.ResponseWriter, r *http.Request) Controller {
	return &real{
		store:     s,
		templates: t,
		writer:    w,
		request:   r,
	}
}

// Get defines a method for filling in the form from the store, if it finds
// nothing then it will return defaults. If an error occurs whilst
// attempting to get, then an error will be rendered.
func (r *real) Get() {
	// Read the store and then render the correct output
	users, err := r.store.Read()
	if err != nil {
		r.render(http.StatusInternalServerError, err)
		return
	}

	if len(users) == 0 {
		r.render(http.StatusNotFound, errors.New("no users found"))
		return
	}

	r.render(http.StatusOK, users)
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

func (r *real) render(code int, data interface{}) {
	r.writer.WriteHeader(code)

	template := r.templates.Get(code)
	template.Execute(r.writer, data)
}
