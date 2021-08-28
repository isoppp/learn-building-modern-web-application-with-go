package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/handlers"
)

const port = ":5555"

func main() {
	handler, err := handlers.NewHandler()
	if err != nil {
		log.Fatal("cannot initialize app")
		return
	}
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
