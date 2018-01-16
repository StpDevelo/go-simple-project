package view

import (
	"bytes"
	"html/template"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	tmplIndex    = parserTemplate("index.tmpl", "app.tmpl", "layout.tmpl")
	tmplNotFound = parserTemplate("not-found.tmpl", "app.tmpl", "layout.tmpl")
)

var (
	m           = minify.New()
	templateDir = "template"
)

type tmpl struct {
	*template.Template
	set []string
}

func init() {

	// add mime types
	mime.AddExtensionType(".js", "text/javascript")

	// add minifier functions
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplateDir(files []string) []string {
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func parserTemplate(set ...string) *tmpl {
	templateName := strings.TrimSuffix(set[0], ".tmpl")
	t := template.New("")
	t.Funcs(template.FuncMap{
		"templateName": func() string {
			return templateName
		},
	})
	_, err := t.ParseFiles(joinTemplateDir(set)...)
	if err != nil {
		log.Fatalf("view: parse template %s error; %v", templateName, err)
	}
	t = t.Lookup("root")
	if t == nil {
		log.Fatalf("view: root template not found in %s", templateName)
	}
	return &tmpl{
		Template: t,
		set:      set,
	}
}

func renderWithStatusCode(w http.ResponseWriter, code int, t *tmpl, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	w.WriteHeader(code)

	pipe := &bytes.Buffer{}

	err := t.Execute(pipe, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = m.Minify("text/html", w, pipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func render(w http.ResponseWriter, t *tmpl, data interface{}) {
	renderWithStatusCode(w, http.StatusOK, t, data)
}
