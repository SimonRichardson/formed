package templates

import "html/template"

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
func NewErrorTemplate() (*template.Template, error) {
	return template.New("error").Parse(errorTemplate)
}

const errorTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Formed - Error!</title>
  </head>
  <body>
    <h1>Error</h1>
    <p>{{.Error}}</p>
  </body>
</html>`

// NewFormTemplate provides a template for the form view
func NewFormTemplate() (*template.Template, error) {
	return template.New("form").Parse(formTemplate)
}

const formTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Formed</title>
  </head>
  <body>
    <form>
		<table>
			<tr>
				<th>First name</th>
				<th>Last name</th>
			</tr>
			{{ range . }}
			<tr>
				<td><input type="text" name="people[][firstname]" value="{{ .FirstName }}" /></td>
				<td><input type="text" name="people[][surname]" value="{{ .Surname }}" /></td>
			</tr>
			{{ end }}
		</table>
		<input type="submit" value="OK" />
	</form>
  </body>
</html>`
