package app

import (
	"net/http"

	"github.com/StpDevelo/go-simple-project/pkg/view"
)

// Mount is Mount Mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", view.Index)
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(&noDir{http.Dir("static")})))
}
