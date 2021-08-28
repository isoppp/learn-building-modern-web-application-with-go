package main

import (
	"fmt"
	"net/http"
)

const port = ":5555"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Home Page")
	if err != nil {
		fmt.Println("Error happens", err)
	}

	fmt.Printf("Bytes written in Home: %d\n", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "About Page")
	if err != nil {
		fmt.Println("Error happens", err)
	}

	fmt.Printf("Bytes written in About: %d\n", n)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server is running on port %s", port)
	_ = http.ListenAndServe(port, nil)
}
