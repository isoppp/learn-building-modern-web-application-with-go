package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/routes"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/renderer"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/controller"
)

const addr = ":5555"

func main() {
	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not initialize template cache")
		return
	}

	c := controller.NewController(tc)
	r := routes.NewRoutes(c)
	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	fmt.Printf("Server is running on port %s\n", addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("Could not initialize app")
		return
	}
}
