package main

import (
	"fmt"
	"net/http"

	"github.com/isoppp/learn-building-modern-web-application-with-go/section3/pkg/handlers"
)

const port = ":5555"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
