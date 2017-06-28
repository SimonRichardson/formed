package controllers

// Controller describes a controller that talks to the underlying models and
// views.
// The controller is envisioned as an interface so that it's possible to
// abstract the API for mocking during testing.
type Controller interface {
	// Get defines a method for filling in the form from the store, if it finds
	// nothing then it will return defaults. If an error occurs whilst
	// attempting to get, then an error will be rendered.
	Get()

	// Post consumes a form that will put the data in to the underlying store.
	// If an error occurs whilst attempting to save, then an error will be
	// rendered.
	Post()

	// NotFound declares a route that doesn't exist, so an error will be
	// rendered.
	NotFound()
}
