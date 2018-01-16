package view

import (
	"net/http"

	"github.com/StpDevelo/go-simple-project/pkg/app"
	"github.com/acoshift/header"
)

type view struct {
	baseURL string
}

// Config is the view config
type Config struct {
	BaseURL string
}

// New Create new view package
func New(config Config) app.View {
	return &view{
		baseURL: config.BaseURL,
	}
}

// Index Render indexPage
func (v *view) Index(w http.ResponseWriter, r *http.Request) {
	render(w, tmplIndex, nil)
}

// NotFound Render not found page
func (v *view) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(header.XContentTypeOptions, "nosniff")
	renderWithStatusCode(w, http.StatusNotFound, tmplNotFound, nil)
}
