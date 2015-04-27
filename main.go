package main

import (
	"net/http"
	"os"
	"github.com/russross/blackfriday"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	// accept a post request and generate the formatted output
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	// serve the public directory
	http.FileServer(http.Dir("public"))
}
