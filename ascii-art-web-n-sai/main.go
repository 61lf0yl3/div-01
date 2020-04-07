package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var err500 string = "500 Internal Server Error"
var err404 string = "404 This page not found"
var err400 string = "400 Bad Request"
var inputFile string = "input.txt"
var inputAscii string = "inputForAscii.txt"
var outputFile string = "output.txt"
var firstRun int = 0

type Page struct {
	Banner string
	Body   []byte
	Output []byte
}

func (p *Page) save() error {
	ioutil.WriteFile(inputFile, p.Body, 0600)
	fmt.Println(p.Body)
	fmt.Println(string(p.Body))
	var enter rune = 13
	var enter2 rune = 10

	res := strings.Replace(string(p.Body), string(enter), "\\", -1)
	res2 := strings.Replace(res, string(enter2), "n", -1)
	fmt.Println(res2)

	return ioutil.WriteFile(inputAscii, []byte(res2), 0600)
}

func loadPage(banner string) (*Page, error) {
	body, err := ioutil.ReadFile(inputFile)
	if firstRun == 0 {
		body = []byte("")
		ioutil.WriteFile(inputFile, []byte(""), 0600)
		ioutil.WriteFile(outputFile, []byte(""), 0600)
	}
	output, err := ioutil.ReadFile(outputFile)
	firstRun = 1
	if err != nil {
		return nil, err
	}
	return &Page{Banner: banner, Body: body, Output: output}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	banner := r.FormValue("banners")
	p, err := loadPage(banner)

	if err != nil {
		fmt.Println(err)
		p = &Page{}
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	banner := r.FormValue("banners")
	p := &Page{Banner: banner, Body: []byte(body)}
	err1 := p.save()
	input, _ := ioutil.ReadFile(inputAscii)
	if !isValid(string(input)) {
		errorHandler(w, r, 400)
		return
	}
	out, err := asciify(string(input), banner)
	if err != nil || err1 != nil {
		errorHandler(w, r, 500)
		return
	}
	ioutil.WriteFile(outputFile, []byte(out), 0600)
	http.Redirect(w, r, "/", http.StatusFound)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	// read file content to buffer
	file.Read(buffer)
	content := bytes.NewReader(buffer)

	// ServeContent uses modtime
	modtime := time.Now()
	// ServeContent uses the name for mime detection
	const name = "ascii.txt"

	// tell the browser the returned content should be downloaded
	w.Header().Add("Content-Disposition", "Attachment")

	http.ServeContent(w, r, name, modtime, content)

	http.Redirect(w, r, "/", http.StatusFound)
}

func isValid(s string) bool {
	for _, letter := range s {
		if letter < 32 || letter > 126 {
			return false
		}
	}
	return true
}

func asciify(args ...string) (string, error) {
	cmd := exec.Command("./ascii", args...)
	cmd.Dir = "ascii"

	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(data))
	return string(data), err
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	t, _ := template.ParseFiles("error.html")
	p := &Page{}
	if status == http.StatusNotFound {
		p = &Page{Banner: err404}
	} else if status == 500 {
		p = &Page{Banner: err500}
	} else if status == 400 {
		p = &Page{Banner: err400}
	}
	t.Execute(w, p)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/download/", downloadHandler)

	fs := http.FileServer(http.Dir("ourCSS"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
