package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
	// "strings"
)

func testing() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/ascii_art", handleAsciiArtWeb)
	http.ListenAndServe(":8080", nil)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {

	//! checks if we are on a wrong "PATH" = `404`
	if r.Method != "/" {
		http.Error(w, "Bad Request", http.StatusNotFound)
		return
	}

	//! checks if we are using the right method in this case which the `GET` request
	if r.Method != "GET" {
		http.Error(w, "Not Found", http.StatusBadRequest)
		return
	}

	//! directing or looking for the file in our project
	tmpl, err := template.ParseFiles("templates/index.html")

	//! if the file does not exist
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//! displays the file to the page
	err = tmpl.Execute(w, nil)

	//! if somthing went wrong while displaying the file
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleAsciiArtWeb(w http.ResponseWriter, r *http.Request) {

	type PageData struct {
		Result string
	}

	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	charMap := loadBanner(banner)

	result := generateAsciiArtWeb(text, charMap)
	data := PageData{Result: result}

	tmpl, err := template.ParseFiles("templates/index.html")

	//! if the file does not exist
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//! displays the file to the page
	err = tmpl.Execute(w, data)

	//! if somthing went wrong while displaying the file
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func loadBannerFile(fileName string) map[rune][]string {
	var styleNames []byte
	var err error

	switch fileName {
	case "shadow.txt":
		styleNames, err = (os.ReadFile("docs/shadow.txt"))
	case "thinkertoy":
		styleNames, err = (os.ReadFile("docs/thinkertoy.txt"))
	default:
		styleNames, err = (os.ReadFile("docs/standard.txt"))
	}
	errorHandlingNew(err)

	byteToString := string(styleNames)
	replaceWithNewLine := strings.ReplaceAll(byteToString, "\r\n", "\n")

	splitContent := strings.Split(replaceWithNewLine, "\n")

	charMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		charLines := splitContent[startLine+1 : startLine+9]
		charMap[rune(i)] = charLines
	}

	return charMap
}

func generateAsciiArtWeb(args string, charMap map[rune][]string) string {
	var result string

	lines := strings.Split(args, "\\n")

	for i, line := range lines{
		
	}
}

func errorHandlingNew(err error) {
	if err != nil {
		panic(err)
	}
}
