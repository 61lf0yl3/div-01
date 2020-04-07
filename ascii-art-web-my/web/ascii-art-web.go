package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Page struct {
	Printtype string
	//Title string
	Input  []byte
	Output []byte
}

func (p *Page) save() error {
	filename := "input.txt"
	return ioutil.WriteFile(filename, p.Input, 0600)
}

func loadPage(printtype string) (*Page, error) {
	filename := "input.txt"
	input, err := ioutil.ReadFile(filename)
	output, err := ioutil.ReadFile("output.txt")
	if err != nil {
		return nil, err
	}
	return &Page{Printtype: printtype, Input: input, Output: output}, nil
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, i love %s", r.URL.Path[1:])
// }
// func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
// 	t, _ := template.ParseFiles(tmpl + ".html")
// 	t.Execute(w, p)
// }
func handler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):] // what difference with r.FormValue
	printtype := r.FormValue("printtype")
	p, err := loadPage(printtype)
	if err != nil {
		fmt.Println(err)
		p = &Page{}
	}
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	renderTemplate(w, "edit", p)
// }

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/save/"):] //why we need it //what difference with r.FormValue  is different thing
	input := r.FormValue("input")                          //take value from .html file with name "input"
	printtype := r.FormValue("printtype")                  //what we give fir input
	p := &Page{Printtype: printtype, Input: []byte(input)} //create node (page) with new parameters
	p.save()
	out := asciify(input[:(len(input))], printtype)     
	ioutil.WriteFile("output.txt", []byte(out), 0600)
	http.Redirect(w, r, "/", http.StatusFound)
	//http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func asciify(args ...string) string {
	cmd := exec.Command("./ascii", args...)
	cmd.Dir = "ascii"

	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(data))
	return string(data)
}

func main() {
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	http.HandleFunc("/", handler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// p1 := &Page{Title: "TestPage", Body: []byte("This is simple Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}
