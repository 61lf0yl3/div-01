package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
}
