package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

// 🇺🇸 User struct for registration
// 🇧🇷 Estrutura de usuário para cadastro
type User struct {
	Name     string
	Email    string
	Password string
	Error    string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

// 🇺🇸 Connects to SQLite using external volume mount
// 🇧🇷 Conecta ao SQLite usando volume externo
func connectDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "/database/users.db")
}

// 🇺🇸 Creates the users table if it doesn't exist
// 🇧🇷 Cria a tabela de usuários se não existir
func createUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

// 🇺🇸 Inserts user with hashed password
// 🇧🇷 Insere usuário com senha criptografada
func insertUser(db *sql.DB, name, email, hashedPassword string) error {
	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, email, hashedPassword)
	return err
}

// 🇺🇸 Handles GET and POST for user registration
// 🇧🇷 Trata GET e POST para cadastro de usuários
func handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	createUserTable(db)

	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "register.html", nil)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := strings.TrimSpace(r.FormValue("password"))

	if name == "" || email == "" || password == "" {
		user := User{Name: name, Email: email, Error: "All fields are required."}
		templates.ExecuteTemplate(w, "register.html", user)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}

	err = insertUser(db, name, email, string(hashedPassword))
	if err != nil {
		user := User{Name: name, Email: email, Error: "Error saving user: email may already be in use."}
		templates.ExecuteTemplate(w, "register.html", user)
		return
	}

	templates.ExecuteTemplate(w, "success.html", name)
}

func main() {
	http.HandleFunc("/user/register", handlerRegisterUser)

	fmt.Println("Server running at http://localhost:8080/user/register")
	http.ListenAndServe(":8080", nil)
}
