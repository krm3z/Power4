package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	path := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "menu", nil)
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "bot", nil)
}

func localHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "local", nil)
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "settings", nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game", nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/bot", botHandler)
	http.HandleFunc("/local", localHandler)
	http.HandleFunc("/settings", settingsHandler)
	http.HandleFunc("/play", gameHandler)

	fmt.Println("✅ Serveur Power4 en écoute sur :8080")
	http.ListenAndServe(":8080", nil)
}
