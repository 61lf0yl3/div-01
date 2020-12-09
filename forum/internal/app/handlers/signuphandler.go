package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// SignupHandler ...
func (c *Connect) SignupHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}
