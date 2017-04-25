package main

import (
	"html/template"
	"net/http"
	"path"
	"./oppgave3"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
	
	speech.SetWitKey("AOW7LJCHDPCGC76I4YTKKWRJDXUVKCUM") //Wit API Key MUST be set before calling any other Wit.AI functions
	speech.SendWitVoice("test.wav")
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