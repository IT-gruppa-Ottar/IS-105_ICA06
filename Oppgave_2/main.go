package main

import (
	"html/template"
	"net/http"
	"path"
)

/**
Generell kode for opprettelse av en html side
 */
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	template, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}