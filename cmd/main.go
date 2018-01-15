package main

import (
	"log"
	"net/http"
	"time"

	"github.com/stpdevelo/go-simple-project/pkg/app"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := http.ListenAndServe(":8000", logger(mux))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func logger(hs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		hs.ServeHTTP(w, r)
		log.Printf("Path is: %s, Methods: %s, Time: %s", r.RequestURI, r.Method, t)
	})
}
