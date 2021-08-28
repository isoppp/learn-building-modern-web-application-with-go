package controller

import (
	"fmt"
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/middlewares"

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
	remoteIP := r.RemoteAddr
	middlewares.GetSession().Put(r.Context(), "remote_ip", remoteIP)
	c.renderer.RenderTemplate(w, "home.page.tmpl")
}

func (c *Controller) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := middlewares.GetSession().GetString(r.Context(), "remote_ip")
	fmt.Println(remoteIP)
	c.renderer.RenderTemplate(w, "about.page.tmpl")
}
