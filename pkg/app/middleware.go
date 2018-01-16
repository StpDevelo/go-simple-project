package app

import (
	"log"
	"net/http"
)

func logger(hs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs.ServeHTTP(w, r)
		log.Printf("Path is: %s, Methods: %s", r.RequestURI, r.Method)
	})
}
