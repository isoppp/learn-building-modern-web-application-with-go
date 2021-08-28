package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0)
	if err != nil {
		_, err2 := fmt.Fprintf(w, "Cannot divide by 0")
		if err2 != nil {
			return
		}
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Value: %f\n", f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)
}
