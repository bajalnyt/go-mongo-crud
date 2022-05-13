package service

import (
	"net/http"
)

// Mux returns a server handler with all the routes.
func Mux() http.Handler {
	apiMux := http.NewServeMux()

	// TODO: set up middleware
	// TODO: set up CORS

	Routes(apiMux)

	return apiMux
}
