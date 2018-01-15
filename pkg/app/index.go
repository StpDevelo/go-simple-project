package app

import (
	"net/http"

	"github.com/stpdevelo/go-simple/pkg/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
