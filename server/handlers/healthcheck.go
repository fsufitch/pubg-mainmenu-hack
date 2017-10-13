package handlers

import "net/http"

type healthCheckHandler struct{}

func (h healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if HSTSRedirect(w, r) {
		return
	}

	w.Write([]byte("OK"))
}
