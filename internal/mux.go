package service

import (
	"net/http"

	"github.com/gobuffalo/logger"
)

// Mux returns a server handler with all the routes.
func Mux(logger logger.Logger) http.Handler {
	apiMux := http.NewServeMux()

	// TODO: set up middleware
	// TODO: set up CORS

	Routes(logger, apiMux)

	return apiMux
}
