package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/renderer"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/handler"
)

const port = ":5555"

func main() {
	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot initialize app")
		return
	}
	h := handler.NewHandler(tc)
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/about", h.About)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
