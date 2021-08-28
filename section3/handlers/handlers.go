package handlers

import (
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/handlers/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.tmpl")
}
