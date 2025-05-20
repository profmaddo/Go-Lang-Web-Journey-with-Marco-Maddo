# LECTURE 05 – SQLite Integration and Basic ORM

## 🧭 Introduction

🇺🇸 In this lecture, we add a critical layer to our Go web project: a database. You will learn how to connect to a **SQLite** database, create tables, insert new records, and retrieve saved data using the Go `database/sql` package. This introduces persistent data storage in your application – a vital step in becoming a full-stack Go developer.

🇧🇷 Nesta aula, adicionamos uma camada crítica ao nosso projeto web com Go: um banco de dados. Você aprenderá como conectar-se a um banco **SQLite**, criar tabelas, inserir novos registros e recuperar dados salvos usando o pacote `database/sql` do Go. Isso introduz armazenamento persistente de dados – um passo essencial para se tornar um desenvolvedor full-stack com Go.

---

## 📄 What You’ll Learn | O que você vai aprender

- ✅ How to use the `database/sql` package with SQLite  
- ✅ How to create tables using SQL (`CREATE TABLE IF NOT EXISTS`)  
- ✅ How to insert data with `Prepare` and `Exec`  
- ✅ How to list and retrieve records using `Query` and `Scan`  
- ✅ Basic data modeling using Go structs

---

## 🧱 CRUD Explained in Go + SQL

| Action | SQL Syntax | Go Function |
|--------|------------|-------------|
| Create | `INSERT INTO` | `stmt.Exec(...)` |
| Read   | `SELECT`       | `db.Query(...)`, `rows.Scan(...)` |
| Update | `UPDATE`       | `stmt.Exec(...) WHERE id=?` |
| Delete | `DELETE`       | `stmt.Exec(...) WHERE id=?` |

---

## 📂 Project Structure

```
/seu-projeto/
├── main.go
├── phones.db                ← database file (auto-created)
├── static/
│   └── styles.css
├── templates/
│   ├── form.html
│   └── list.html
```

---

## 🚀 Run and Test

```bash
go run main.go
```

Access the app:

- `http://localhost:8080/register` → Register a new phone
- `http://localhost:8080/phones`   → List all registered phones

---

## 🧠 Integration with Previous Lectures

- ✅ HTML Templates: still uses `html/template` (from Lecture 2)  
- ✅ CSS Styling: served from `/static/` (from Lecture 3)  
- ✅ Form Handling: validation logic with `r.FormValue()` (from Lecture 4)

Now we persist form data using **SQLite**, expanding the application to be stateful and closer to real-world use.

---

## 🔍 About SQLite in Go

- Uses file-based storage (lightweight `.db` file)
- Requires import of driver: `_ "github.com/mattn/go-sqlite3"`
- Queries and executions use standard SQL syntax
- Ideal for learning full-stack development

---

Happy coding! | Bons estudos!

---

## 🧩 Installing `github.com/mattn/go-sqlite3`

🇺🇸 To use SQLite in your Go project, you must import the following driver:

```go
_ "github.com/mattn/go-sqlite3"
```

Then, in the project root folder:

```bash
go mod init nome_do_modulo   # only once
go get github.com/mattn/go-sqlite3
```

🇧🇷 Para usar SQLite em seu projeto Go, você precisa importar o seguinte driver:

```go
_ "github.com/mattn/go-sqlite3"
```

E no terminal, na pasta do seu projeto:

```bash
go mod init nome_do_modulo   # apenas uma vez
go get github.com/mattn/go-sqlite3
```

---

## ⚠️ CGO Enabled Required

🇺🇸 This package relies on **CGO**, which means you must compile with `CGO_ENABLED=1`.  
If you're building in Docker, make sure **NOT** to disable CGO.  
Remove `CGO_ENABLED=0` from your Dockerfile.

🇧🇷 Este pacote depende do **CGO**, ou seja, você precisa compilar com `CGO_ENABLED=1`.  
Se estiver usando Docker, certifique-se de **NÃO** desativar o CGO.  
Remova `CGO_ENABLED=0` do seu Dockerfile.

---

## 💡 Why this step is important?

🇺🇸 Without installing the driver and enabling CGO, your project will not compile.  
SQLite will not be accessible, and the `import` will return an error.

🇧🇷 Sem instalar o driver e ativar o CGO, seu projeto não compilará.  
O SQLite não estará disponível e o `import` causará erro.

