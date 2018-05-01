package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// IndexPageData represents the data to be rendered by the index page template
type IndexPageData struct {
	Message  string
	Title    string
	Hostname string
}

// HandleIndex renders the index page for the web server
func HandleIndex(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	data := &IndexPageData{
		Message:  "Hello World",
		Title:    "Index Page",
		Hostname: hostname,
	}

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := tpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
