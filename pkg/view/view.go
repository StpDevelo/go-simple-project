package view

import (
	"net/http"

	"github.com/acoshift/header"
)

// Index Render indexPage
func Index(w http.ResponseWriter, r *http.Request) {
	render(w, tmplIndex, nil)
}

// NotFound Render not found page
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(header.XContentTypeOptions, "nosniff")
	renderWithStatusCode(w, http.StatusNotFound, tmplNotFound, nil)
}
