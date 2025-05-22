package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func connectDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "/database/users.db")
}

func validateLogin(email, password string) bool {
	db, err := connectDB()
	if err != nil {
		return false
	}
	defer db.Close()

	var hashed string
	row := db.QueryRow("SELECT password FROM users WHERE email = ?", email)
	err = row.Scan(&hashed)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_email")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "login.html", nil)
	case "POST":
		email := strings.TrimSpace(r.FormValue("email"))
		password := strings.TrimSpace(r.FormValue("password"))

		if validateLogin(email, password) {
			cookie := http.Cookie{
				Name:    "session_email",
				Value:   email,
				Expires: time.Now().Add(1 * time.Hour),
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			templates.ExecuteTemplate(w, "error.html", "Invalid email or password.")
		}
	}
}

func handlerDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_email")
	email := cookie.Value
	templates.ExecuteTemplate(w, "dashboard.html", email)
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "session_email",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/dashboard", authMiddleware(handlerDashboard))
	http.HandleFunc("/logout", handlerLogout)

	fmt.Println("Server running at http://localhost:8080/login")
	http.ListenAndServe(":8080", nil)
}
