package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StpDevelo/go-simple-project/pkg/app"
	"github.com/StpDevelo/go-simple-project/pkg/controller"
	"github.com/StpDevelo/go-simple-project/pkg/view"
	"github.com/acoshift/configfile"
)

func main() {

	config := configfile.NewReader("config")

	view := view.New(view.Config{
		BaseURL: config.String("base_url"),
	})
	ctrl := controller.New(controller.Config{
		View: view,
	})

	app := app.New(app.Config{
		Controller: ctrl,
		View:       view,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	mux.Handle("/", app)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
