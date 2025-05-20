package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// 🇺🇸 Struct representing a phone submitted by the user
// 🇧🇷 Estrutura que representa um celular enviado pelo usuário
type Phone struct {
	Model string
	Brand string
	Price string
	Error string
}

// 🇺🇸 Templates global variable
// 🇧🇷 Variável global para os templates
var templates = template.Must(template.ParseGlob("templates/*.html"))

// 🇺🇸 Renders the form page
// 🇧🇷 Renderiza a página com o formulário
func handlerForm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// 🇺🇸 Show empty form
		// 🇧🇷 Mostra o formulário vazio
		templates.ExecuteTemplate(w, "form.html", nil)
	case "POST":
		// 🇺🇸 Get form values
		// 🇧🇷 Captura valores do formulário
		model := strings.TrimSpace(r.FormValue("model"))
		brand := strings.TrimSpace(r.FormValue("brand"))
		price := strings.TrimSpace(r.FormValue("price"))

		// 🇺🇸 Validate inputs
		// 🇧🇷 Valida entradas
		if model == "" || brand == "" || price == "" {
			phone := Phone{Model: model, Brand: brand, Price: price, Error: "All fields are required."}
			templates.ExecuteTemplate(w, "form.html", phone)
			return
		}

		// 🇺🇸 Show success message
		// 🇧🇷 Mostra mensagem de sucesso
		phone := Phone{Model: model, Brand: brand, Price: price}
		templates.ExecuteTemplate(w, "success.html", phone)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/register", handlerForm)

	fmt.Println("Server running at http://localhost:8080/register")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
