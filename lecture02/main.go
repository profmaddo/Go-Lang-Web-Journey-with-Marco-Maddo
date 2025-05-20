package main

import (
	"html/template"
	"net/http"
	"fmt"
)

// 🇺🇸 Phone represents a sample struct to render data in HTML
// 🇧🇷 Phone representa uma estrutura de exemplo para renderizar dados no HTML
type Phone struct {
	Model string
	Brand string
	Price float64
}

// 🇺🇸 Templates variable to parse all HTML files in the templates folder
// 🇧🇷 Variável templates para carregar todos os arquivos HTML da pasta templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

// 🇺🇸 Home handler displays a list of phones using a template
// 🇧🇷 Handler da página inicial exibe uma lista de celulares usando um template
func handlerHome(w http.ResponseWriter, r *http.Request) {
	phones := []Phone{
		{"iPhone 13", "Apple", 4999.90},
		{"Galaxy S22", "Samsung", 3999.00},
		{"Moto G100", "Motorola", 2299.50},
	}

	err := templates.ExecuteTemplate(w, "index.html", phones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 🇺🇸 About page handler
// 🇧🇷 Handler da página "Sobre"
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/about", handlerAbout)

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
