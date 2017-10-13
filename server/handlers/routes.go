package handlers

import "github.com/gorilla/mux"

// RegisterHandlers registers all the handlers at appropriate routes
func RegisterHandlers(router *mux.Router) {
	// Public endpoints
	router.Path("/health").Handler(healthCheckHandler{})
}

// RegisterEvilIndexHandler registers the special "evil" endpoint to serve a fake index
func RegisterEvilIndexHandler(router *mux.Router, realPUBGURL string, staticPath string, apiHost string) error {
	handler, err := NewEvilIndexHandler(realPUBGURL, staticPath, apiHost)
	if err != nil {
		return err
	}
	router.Path("/index.html").Handler(handler)
	return nil
}
