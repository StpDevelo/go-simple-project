package app

import "net/http"

// View is an app view render
type View interface {
	Index(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
}
