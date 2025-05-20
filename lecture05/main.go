package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// 🇺🇸 Struct for phone model
// 🇧🇷 Estrutura para representar um celular
type Phone struct {
	ID    int
	Model string
	Brand string
	Price string
}

// 🇺🇸 Template parser
// 🇧🇷 Carregador de templates HTML
var templates = template.Must(template.ParseGlob("templates/*.html"))

// 🇺🇸 Opens database connection
// 🇧🇷 Abre a conexão com o banco de dados SQLite
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "phones.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 🇺🇸 Creates the phone table if not exists
// 🇧🇷 Cria a tabela "phones" se não existir
func createTableIfNotExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS phones (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		model TEXT,
		brand TEXT,
		price TEXT
	);
	`
	_, err := db.Exec(query)
	return err
}

// 🇺🇸 Inserts new phone into the database
// 🇧🇷 Insere novo celular no banco de dados
func insertPhone(db *sql.DB, model, brand, price string) error {
	stmt, err := db.Prepare("INSERT INTO phones (model, brand, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(model, brand, price)
	return err
}

// 🇺🇸 Lists all phones from the database
// 🇧🇷 Lista todos os celulares cadastrados
func listPhones(db *sql.DB) ([]Phone, error) {
	rows, err := db.Query("SELECT id, model, brand, price FROM phones")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phones []Phone
	for rows.Next() {
		var p Phone
		rows.Scan(&p.ID, &p.Model, &p.Brand, &p.Price)
		phones = append(phones, p)
	}
	return phones, nil
}

// 🇺🇸 Handles phone registration
// 🇧🇷 Trata o registro de celulares via formulário
func handlerRegister(w http.ResponseWriter, r *http.Request) {
	db, _ := connectDB()
	createTableIfNotExists(db)

	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "form.html", nil)
	case "POST":
		model := strings.TrimSpace(r.FormValue("model"))
		brand := strings.TrimSpace(r.FormValue("brand"))
		price := strings.TrimSpace(r.FormValue("price"))

		if model == "" || brand == "" || price == "" {
			templates.ExecuteTemplate(w, "form.html", "All fields are required")
			return
		}

		err := insertPhone(db, model, brand, price)
		if err != nil {
			http.Error(w, "Error saving data", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/phones", http.StatusSeeOther)
	}
}

// 🇺🇸 Shows all phones
// 🇧🇷 Exibe todos os celulares salvos
func handlerList(w http.ResponseWriter, r *http.Request) {
	db, _ := connectDB()
	createTableIfNotExists(db)

	phones, err := listPhones(db)
	if err != nil {
		http.Error(w, "Error loading data", http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "list.html", phones)
}

func main() {
	http.HandleFunc("/register", handlerRegister)
	http.HandleFunc("/phones", handlerList)

	fmt.Println("Server running at http://localhost:8080/register")
	http.ListenAndServe(":8080", nil)
}
