package main

import (
	"fmt"
	"net/http"
	"os"
)

// handlerHome lê e exibe o conteúdo de index.html
func handlerHome(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Unable to load homepage", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// handlerAbout responde quando o usuário acessa "/about"
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a basic web server built in Go.")
}

func main() {
	// Associa os caminhos "/" e "/about" às funções específicas
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/about", handlerAbout)

	// Inicia o servidor na porta 8080 e imprime log no terminal
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // nil usa o DefaultServeMux
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
