package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type templateCache map[string]*template.Template

type Handler struct {
	templateCache
}

func NewHandler() (*Handler, error) {
	var handler Handler
	tc, err := createTemplateCache()
	if err != nil {
		return &handler, err
	}

	handler.templateCache = tc
	return &handler, err
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "home.page.tmpl")
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "about.page.tmpl")
}

func (h *Handler) renderTemplate(w http.ResponseWriter, tmpl string) {
	t, ok := h.templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get tempalte from cache")
		return
	}

	err := t.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template : ", err)
		return
	}
}

func createTemplateCache() (templateCache, error) {
	functions := template.FuncMap{}
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./section3/templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./section3/templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./section3/templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
