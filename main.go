package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
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

	if banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy.txt" {
		http.Error(w, "Not Found", http.StatusNotFound) // 404
		return
	}

	charMap := loadBanner(banner)

	result := generateAsciiArt(text, charMap)

	data := PageData{Result: result}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	//! checks for wrong PATH is `404`
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//! checks for wrong METHOD is `400`
	if r.Method != "GET" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	//! checks for error while LOADING the page `404`
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)

	//! checks for INTERNAL-SERVER_ERROR `500`
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func loadBanner(fileName string) map[rune][]string {
	var readingByte []byte // ← declare outside
	var err error

	//^ reading the file
	switch fileName {
	case "shadow.txt":
		readingByte, err = (os.ReadFile("docs/shadow.txt"))
	case "thinkertoy.txt":
		readingByte, err = (os.ReadFile("docs/thinkertoy.txt"))
	default:
		readingByte, err = (os.ReadFile("docs/standard.txt"))

	}
	errorHandling(err)

	//^ converting byte to string
	convertToString := string(readingByte)
	convertToString = strings.ReplaceAll(convertToString, "\r\n", "\n")

	//! split the strings using NEWLINE
	splitStr := strings.Split(convertToString, "\n")

	//^ create a map for the LETTERS or ASCII
	charMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		charLines := splitStr[startLine+1 : startLine+9]
		charMap[rune(i)] = charLines
	}

	return charMap
}

func generateAsciiArt(args string, charMap map[rune][]string) string {
	var result string

	lines := strings.Split(args, "\\n")

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				result += "\n"
			}
			continue
		}
		row := make([]string, 8)

		for _, char := range line {
			for i := 0; i < 8; i++ {
				row[i] += charMap[char][i]
			}
		}
		for _, r := range row {
			result += r + "\n"
		}
	}
	return result
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleHome)
    http.HandleFunc("/ascii-art", handleAsciiArt)
    http.ListenAndServe(":8080", nil)
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
