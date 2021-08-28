package routes

import (
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/middlewares"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRoutes(c *controller.Controller) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", c.Home)
	mux.Get("/about", c.About)
	return mux
}
