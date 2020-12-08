package server

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/61lf0yl3/div-01/forum/internal/app/database"
)

// Server ...
type Server struct {
	database *database.Database
}

// Start ...
func Start() error {

	server := Server{}

	if err := server.database.InitDB("forum.db"); err != nil {
		return err
	}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/usercreate", userCreateHandler)
	return http.ListenAndServe(":8001", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	//	email := r.FormValue("email")

}
