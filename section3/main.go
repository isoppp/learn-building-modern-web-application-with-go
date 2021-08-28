package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":5555"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err1 := template.ParseFiles("./section3/templates/" + tmpl)

	if err1 != nil {
		fmt.Println("Error in ParseFiles : ", err1)
		return
	}

	err2 := parsedTemplate.Execute(w, nil)

	if err2 != nil {
		fmt.Println("Error parsing template : ", err2)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server is running on port %s\n", port)
	_ = http.ListenAndServe(port, nil)

}
