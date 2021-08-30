package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/front/home/index.html", "views/front/layout.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "layout", nil)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/front/news/index.html", "views/front/layout.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "layout", nil)
}

func main() {

	fmt.Println("Lisetning ...")
	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/yangiliklar", newsHandler)

	http.ListenAndServe("localhost:8001", nil)
}