package app

import (
	"net/http"
)

// Mount is Mount Mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(&noDir{http.Dir("static")})))
}
