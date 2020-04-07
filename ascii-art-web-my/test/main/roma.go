package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	//errorHandler(w, r, 404)
	// 	w.Write([]byte("<h1>404 Oops, this page not found...</h1>"))
	// 	return
	// }
	//tmplt, err := template.ParseFiles("index2.html")
	//if err != nil {
	// 	fmt.Println(err)
	// }
	// tmplt.Execute(w, page)
	w.Write([]byte("<h1>РОМА ПИЗДАБОЛ!!!</h1>"))
}

func main() {
	http.HandleFunc("/", indexHandler)

	//fs := http.FileServer(http.Dir("../css"))
	//http.Handle("/css/", http.StripPrefix("/css/", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
