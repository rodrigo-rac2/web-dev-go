package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderHtml(w http.ResponseWriter, htmlTemplate string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + htmlTemplate)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}
