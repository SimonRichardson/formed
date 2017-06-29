package controllers

import (
	"net/http"

	"github.com/SimonRichardson/formed/pkg/store"
	"github.com/SimonRichardson/formed/pkg/templates"
	"github.com/pkg/errors"
)

// go doesn't provide this sort of form iteration, so it just provides a slice
const (
	formKeyFirstName = "people[][firstname]"
	formKeySurname   = "people[][surname]"
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
	if err := r.request.ParseForm(); err != nil {
		r.render(http.StatusBadRequest, errors.Wrap(err, "invalid form data"))
		return
	}

	// Extract the firstnames, surnames
	var userForm UserForm
	if err := userForm.DecodeFrom(r.request.Form); err != nil {
		r.render(http.StatusBadRequest, errors.Wrap(err, "invalid form user data"))
		return
	}

	// Convert the form data to actual users
	users, err := userForm.Users()
	if err != nil {
		r.render(http.StatusBadRequest, errors.Wrap(err, "invalid user data"))
		return
	}

	// Write the users to the underlying store
	if err := r.store.Write(users); err != nil {
		r.render(http.StatusInternalServerError, errors.Wrap(err, "invalid user data"))
		return
	}

	// Once we've written, let's redirect to the correct page
	http.Redirect(r.writer, r.request, "/query", http.StatusSeeOther)
}

// NotFound declares a route that doesn't exist, so an error will be
// rendered.
func (r *real) NotFound() {
	r.render(http.StatusNotFound, errors.New("not found"))
}

func (r *real) render(code int, data interface{}) {
	r.writer.WriteHeader(code)

	template := r.templates.Get(code)
	template.Execute(r.writer, data)
}
