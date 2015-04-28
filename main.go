package main

import (
	"net/http"
	"os"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	// accept a post request and generate the formatted output
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	html := bluemonday.UGCPolicy().SanitizeBytes(markdown)
	rw.Write(html)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	// serve the public directory
	http.FileServer(http.Dir("public"))
}
