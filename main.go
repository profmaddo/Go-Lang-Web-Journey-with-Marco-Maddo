package main

import (
	"html/template"
	"log"
	"net/http"
)

// Estrutura do item de menu
type Lecture struct {
	Title string
	Path  string
}

func main() {
	http.HandleFunc("/", homeHandler)

	// Handlers para cada lecture – você criará cada um em seu respectivo pacote/module
	http.HandleFunc("/lecture01", lecture01Handler)
	http.HandleFunc("/lecture02", lecture02Handler)
	http.HandleFunc("/lecture03", lecture03Handler)
	http.HandleFunc("/lecture04", lecture04Handler)
	http.HandleFunc("/lecture05", lecture05Handler)
	http.HandleFunc("/lecture06", lecture06Handler)
	http.HandleFunc("/lecture07", lecture07Handler)

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lectures := []Lecture{
		{"LECTURE 01 – Getting Started with Web Development in Go", "/lecture01"},
		{"LECTURE 02 – HTML Templating with Go", "/lecture02"},
		{"LECTURE 03 – Serving Static Files in Go", "/lecture03"},
		{"LECTURE 04 – Handling Forms and Validations in Go", "/lecture04"},
		{"LECTURE 05 – SQLite Integration in Go", "/lecture05"},
		{"LECTURE 06 – CRUD Operations with SQLite and Go", "/lecture06"},
		{"LECTURE 07 – User Authentication and Session Management in Go", "/lecture07"},
	}

	tmpl := template.Must(template.New("home").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Go Web Course – Main Menu</title>
			<style>
				body { font-family: Arial, sans-serif; padding: 2rem; }
				h1 { color: #2c3e50; }
				ul { list-style: none; padding: 0; }
				li { margin-bottom: 10px; }
				a { text-decoration: none; color: #2980b9; }
				a:hover { text-decoration: underline; }
			</style>
		</head>
		<body>
			<h1>Go Web Course – Main Menu</h1>
			<ul>
				{{range .}}
					<li><a href="{{.Path}}">{{.Title}}</a></li>
				{{end}}
			</ul>
		</body>
		</html>
	`))

	err := tmpl.Execute(w, lectures)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funções de stub (você pode apontar para os pacotes de cada lecture posteriormente)
func lecture01Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 01 - Conteúdo aqui..."))
}
func lecture02Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 02 - Conteúdo aqui..."))
}
func lecture03Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 03 - Conteúdo aqui..."))
}
func lecture04Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 04 - Conteúdo aqui..."))
}
func lecture05Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 05 - Conteúdo aqui..."))
}
func lecture06Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 06 - Conteúdo aqui..."))
}
func lecture07Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lecture 07 - Conteúdo aqui...."))
}
