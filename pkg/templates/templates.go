package templates

import (
	"html/template"

	"github.com/pkg/errors"
)

// Templates holds a key, value store of templates that can be used to render
// depending on the key required.
// Note: it also has a fallback template if nothing if found so that we can
// display a valid error to the user.
type Templates struct {
	templates map[int]*template.Template
	fallback  *template.Template
}

// NewTemplates creates a Template key, value store with an additional fallback
func NewTemplates(fallback *template.Template) *Templates {
	return &Templates{
		templates: make(map[int]*template.Template),
		fallback:  fallback,
	}
}

// Get returns a template depending on the key supplied, otherwise it will
// return the fallback
func (t *Templates) Get(key int) *template.Template {
	if t, ok := t.templates[key]; ok {
		return t
	}
	return t.fallback
}

// Set provides a way to set a template for a specific key
func (t *Templates) Set(key int, tmpl *template.Template) {
	t.templates[key] = tmpl
}

// NewErrorTemplate provides a template for all generic errors
func NewErrorTemplate(useLocal bool) (*template.Template, error) {
	tmpl, err := FSString(useLocal, "/views/error.html")
	if err != nil {
		return nil, errors.Wrap(err, "unable to load template")
	}
	return template.New("error").Parse(tmpl)
}

// NewFormTemplate provides a template for the form view
func NewFormTemplate(useLocal bool) (*template.Template, error) {
	tmpl, err := FSString(useLocal, "/views/index.html")
	if err != nil {
		return nil, errors.Wrap(err, "unable to load template")
	}
	return template.New("form").Parse(tmpl)
}
