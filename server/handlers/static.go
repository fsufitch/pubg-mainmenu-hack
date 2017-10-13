package handlers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// StaticHandler is a handler to serve cached static bytes
type StaticHandler struct {
	data []byte
	etag string
}

// NewStaticHandler creates a handler to serve cached static bytes
func NewStaticHandler(data []byte) *StaticHandler {
	h := StaticHandler{}
	h.data = data

	sum := md5.Sum(data)
	h.etag = fmt.Sprintf(`"%s"`, base64.StdEncoding.EncodeToString(sum[:]))

	return &h
}

func (h StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Etag", h.etag)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, h.etag) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(h.data)
}
