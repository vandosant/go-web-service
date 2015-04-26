package main

import (
  "net/http"
  "github.com/russross/blackfriday"
  "fmt"
)

func main() {
  http.HandleFunc("/markdown", GenerateMarkdown)
  http.Handle("/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
  fmt.Println(r.FormValue("body"))
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}
