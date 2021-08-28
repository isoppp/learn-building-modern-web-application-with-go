package renderer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

type Renderer struct {
	templateCache TemplateCache
}

func NewRenderer(tc TemplateCache) *Renderer {
	var r Renderer
	r.templateCache = tc
	return &r
}

func (h *Renderer) RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, ok := h.templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get tempalte from cache")
		return
	}

	data := &struct {
		Static struct {
			SiteName string
		}
	}{
		Static: struct{ SiteName string }{SiteName: "Example Site Name passed by render"},
	}
	err := t.Execute(w, data)

	if err != nil {
		fmt.Println("Error parsing template : ", err)
		return
	}
}

func CreateTemplateCache() (TemplateCache, error) {
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
