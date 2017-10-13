package handlers

import (
	"net/http"
	"net/url"
)

func writeClientError(w http.ResponseWriter, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte("client error: " + text))
}

func writeServerError(w http.ResponseWriter, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte("server error: " + text))
}

// HerokuSSLRedirectHost holds the host where HTTPS redirection should go to
var HerokuSSLRedirectHost = ""

const hstsHeaderValue = "max-age=18000; includeSubDomains" // 5 minutes

// HSTSRedirect redirects to using SSL, if configured to do so
// If redirected, it also sets HSTS headers.
// Returns true if redirection happened
func HSTSRedirect(w http.ResponseWriter, r *http.Request) bool {
	if HerokuSSLRedirectHost == "" {
		return false
	}

	if r.Header.Get("X-Forwarded-Proto") == "http" {
		url, _ := url.Parse(r.URL.String())

		url.Scheme = "https"
		url.Host = HerokuSSLRedirectHost

		w.Header().Set("Location", url.String())
		w.WriteHeader(301)
		w.Write([]byte(""))
		return true
	}

	w.Header().Set("Strict-Transport-Security", hstsHeaderValue)
	return false
}
