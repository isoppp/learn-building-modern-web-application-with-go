package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
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
