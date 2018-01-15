package app

import (
	"net/http"

	"github.com/stpdevelo/go-simple-project/pkg/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.TmplIndex, nil)
}
