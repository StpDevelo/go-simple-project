package app

import "net/http"

// Controller is an app controller
type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)
}
