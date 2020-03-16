package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	ascii ".."
)

var err500 = "500 Internal Server Error :("
var err404 = "404 Oops, this page not found..."
var err400 = "400 Bad request, I can't print this!"
var outputFile string = "outpit.txt"

type Page struct {
	Body   string
	Font   string
	Output string
	//Error  []byte
}

var page Page

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//errorHandler(w, r, 404)
		w.Write([]byte("<h1>404 Oops, this page not found...</h1>"))
		return
	}
	tmplt, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmplt.Execute(w, page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	body = strings.ReplaceAll(body, "\r\n", "\\n")
	fmt.Println("this is body", body)
	if !isValid(body) {
		//errorHandler(w, r, 404)
		w.Write([]byte("<h1>400 Bad request, I can't print this!</h1>"))
		return
	}
	page.Font = r.FormValue("banners")
	fmt.Println("this is font", page.Font)
	page.Body = body
	asciiarr := []string{body, page.Font}
	fmt.Println("this is asciiarray", asciiarr)
	out, err := ascii.Asciify(asciiarr)
	if err != nil {
		fmt.Println("Error in asciify", err) //internal error in server 500
		//errorHandler(w, r, 500)
		w.Write([]byte("<h1>500 Internal Server Error :(</h1>"))
		return
	}
	page.Output = string(out)
	fmt.Println("this is output", out)
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

	w.Header().Set("Content-Type", "multipart/form-data; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii.txt")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(int(size)))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	http.ServeContent(w, r, "ascii.txt", time.Now(), bytes.NewReader(buffer))
	http.Redirect(w, r, "/", http.StatusFound)
}

func isValid(str string) bool {
	for _, l := range str {
		if l < 32 || l > 126 {
			return false
		}
	}
	return true
}

// func errorHandler(w http.ResponseWriter, r *http.Request, err int) {

// }

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/download/", downloadHandler)

	fs := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
