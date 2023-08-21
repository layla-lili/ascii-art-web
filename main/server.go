package main

import (
	"asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("indexF.html"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Handler for the root URL ("/")
	http.HandleFunc("/", homeHandler)
	// Handler for the "/ascii-art" URL
	http.HandleFunc("/ascii-art", asciiArtHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// homeHandler handles GET requests to the root URL ("/").
// It renders the indexF.html template.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "indexF.html", nil)
	}
}

// asciiArtHandler handles POST requests to the "/ascii-art" URL.
// It generates ASCII art based on the input text and selected banner,
// and renders the indexF.html template with the generated ASCII art.

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		asciiArt := generateAsciiArt(text, banner)
		templates.ExecuteTemplate(w, "indexF.html", struct {
			Text     string
			Banner   string
			AsciiArt string
		}{
			Text:     text,
			Banner:   banner,
			AsciiArt: asciiArt,
		})
	}
}

func generateAsciiArt(text, banner string) string {
	// Implement your ASCII art generation logic based on the selected banner
	// Here's a simple example for the three banners mentioned
	file, err := os.Open(banner+".txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
	return 	asciiart.ReadLine(text, file)

}
