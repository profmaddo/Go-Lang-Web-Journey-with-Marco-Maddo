package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 🇺🇸 Struct representing a phone model
// 🇧🇷 Estrutura que representa um modelo de celular
type Phone struct {
	Model string
	Brand string
	Price float64
}

// 🇺🇸 Parse all templates from templates folder
// 🇧🇷 Analisa todos os templates da pasta templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

// 🇺🇸 Renders the home page with a list of phones
// 🇧🇷 Renderiza a página inicial com uma lista de celulares
func handlerHome(w http.ResponseWriter, r *http.Request) {
	phones := []Phone{
		{"iPhone 14", "Apple", 5999.99},
		{"Pixel 7", "Google", 4799.00},
		{"Galaxy A73", "Samsung", 2399.50},
	}
	err := templates.ExecuteTemplate(w, "index.html", phones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 🇺🇸 Renders a simple about page
// 🇧🇷 Renderiza uma página simples "Sobre"
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 🇺🇸 Serve static files from /static/ directory
	// 🇧🇷 Serve arquivos estáticos da pasta /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 🇺🇸 Define route handlers
	// 🇧🇷 Define os manipuladores de rota
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/about", handlerAbout)

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
