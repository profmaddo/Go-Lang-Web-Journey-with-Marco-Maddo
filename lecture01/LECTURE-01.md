# LECTURE 01 – Getting Started with Web Development in Go

## 🧭 Introduction

In this first lecture, we will learn how to create a basic web server using the Go programming language. This is the foundation for any web development project in Go. We’ll introduce the `net/http` package, which is the standard tool provided by the Go standard library to handle web servers and routes. You’ll create a program that responds to browser requests using two simple pages: the homepage (`/`) and an about page (`/about`). This lecture sets the stage for understanding request-response cycles, routing, and Go’s built-in web capabilities with zero external dependencies.

Nesta primeira aula, aprenderemos como criar um servidor web básico usando a linguagem de programação Go. Esta é a base para qualquer projeto de desenvolvimento web em Go. Apresentaremos o pacote `net/http`, que é a ferramenta padrão fornecida pela biblioteca padrão da linguagem para lidar com servidores web e rotas. Você criará um programa que responde a requisições do navegador usando duas páginas simples: a página inicial (`/`) e uma página "sobre" (`/about`). Esta aula define os fundamentos para compreender o ciclo de requisição e resposta, roteamento, e as capacidades web integradas do Go sem depender de bibliotecas externas.

---

## ✅ Code Summary 

```go
http.HandleFunc("/", handlerHome)
```
🇺🇸 Maps the route `/` to the `handlerHome` function.  
🇧🇷 Mapeia a rota `/` para a função `handlerHome`.

```go
http.ListenAndServe(":8080", nil)
```
🇺🇸 Starts the server on port 8080 using the default router.  
🇧🇷 Inicia o servidor na porta 8080 usando o roteador padrão.

```go
func handlerHome(w http.ResponseWriter, r *http.Request)
```
🇺🇸 Handles the homepage request and writes a welcome message.  
🇧🇷 Trata o acesso à página inicial e escreve uma mensagem de boas-vindas.

```go
func handlerAbout(w http.ResponseWriter, r *http.Request)
```
🇺🇸 Handles the `/about` route with a description.  
🇧🇷 Trata a rota `/about` com uma descrição sobre o projeto.

---

## 🌐 How to Run

1. Save the file as `main.go`
2. Open terminal in the same directory
3. Run:
```bash
go run main.go
```
4. Open your browser and access :
```
http://localhost:8080/
http://localhost:8080/about
```

---

Happy coding! | Bons estudos!
