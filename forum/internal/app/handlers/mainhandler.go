package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// MainHandler ...
func (c *Connect) MainHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}
