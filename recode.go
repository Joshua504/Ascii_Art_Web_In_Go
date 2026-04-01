package main

import (
	"html/template"
	"net/http"
)

func testing() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/ascii_art", handleAsciiArtWeb)
	http.ListenAndServe(":8080", nil)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")

	if err != nil{
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func handleAsciiArtWeb(w http.ResponseWriter, r *http.Request){
	
}