package controller

import (
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/renderer"
)

type Controller struct {
	renderer *renderer.Renderer
}

func NewController(tc renderer.TemplateCache) *Controller {
	var controller Controller
	controller.renderer = renderer.NewRenderer(tc)
	return &controller
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	c.renderer.RenderTemplate(w, "home.page.tmpl")
}

func (c *Controller) About(w http.ResponseWriter, r *http.Request) {
	c.renderer.RenderTemplate(w, "about.page.tmpl")
}
