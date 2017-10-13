package resources

import "net/http"

const redirectURL = "https://github.com/fsufitch/pubg-mainmenu-hack"

// FallbackHandler is a redirect to this project's Github page for documentation purposes
type FallbackHandler struct{}

func (h FallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", redirectURL)
	w.WriteHeader(http.StatusPermanentRedirect)
}
