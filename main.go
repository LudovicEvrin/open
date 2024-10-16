package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	// Servir les fichiers statiques (JavaScript, CSS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Définir les routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)

	// Lancer le serveur
	log.Println("Serveur démarré sur : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	num1Str := r.URL.Query().Get("num1")
	num2Str := r.URL.Query().Get("num2")

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		http.Error(w, "Attention ! Les entrées doivent être des nombres valides.", http.StatusBadRequest)
		return
	}

	result := num1 + num2
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"result": %d}`, result)
}
