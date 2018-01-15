package view

import (
	"html/template"
	"net/http"
)

var (
	tpIndex = template.New("")
)

func init() {
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")
	if err != nil {
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}

// Index is render func for view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
