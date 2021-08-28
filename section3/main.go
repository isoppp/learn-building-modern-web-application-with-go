package main

import (
	"fmt"
	"net/http"
)

const port = ":5555"

func Home(w http.ResponseWriter, r *http.Request) {
}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
