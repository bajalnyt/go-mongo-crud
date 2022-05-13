package service

import "net/http"

// Routes sets up the routes.
func Routes(api *http.ServeMux) {

	// healthcheck
	api.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

}
