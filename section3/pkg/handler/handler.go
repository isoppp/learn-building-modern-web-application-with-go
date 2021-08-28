package handler

import (
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/renderer"
)

type Handler struct {
	renderer *renderer.Renderer
}

func NewHandler(tc renderer.TemplateCache) *Handler {
	var handler Handler
	handler.renderer = renderer.NewRenderer(tc)
	return &handler
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "home.page.tmpl")
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderTemplate(w, "about.page.tmpl")
}
