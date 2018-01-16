package app

import (
	"net/http"

	"github.com/acoshift/middleware"
)

type app struct {
	http.Handler
	view View
	ctrl Controller
}

// New create new app
func New(config Config) http.Handler {
	view := config.View
	ctrl := config.Controller
	app := &app{
		view: view,
		ctrl: ctrl,
	}

	mux := http.NewServeMux()
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(&noDir{http.Dir("static")})))

	main := http.NewServeMux()
	main.Handle("/", http.HandlerFunc(ctrl.Index))

	mux.Handle("/", main)

	app.Handler = middleware.Chain(
		logger,
	)(mux)

	return app
}
