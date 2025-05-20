# LECTURE B01 – Docker Setup for Go Web App

## 🧭 Introduction

In this optional but highly useful lecture, we’ll create a Docker environment to run your Go web application. Docker allows you to create isolated containers that work the same in every machine, making local development and future deployment much easier. Starting with this lecture, each subsequent part of the project will be compatible with Docker.

Nesta aula opcional porém extremamente útil, criaremos um ambiente Docker para executar sua aplicação web escrita em Go. O Docker permite criar containers isolados que funcionam da mesma forma em qualquer máquina, facilitando o desenvolvimento local e o deploy futuro. A partir desta aula, cada parte do projeto será compatível com Docker.

---

## ❗ Common Docker Build Error

### ❌ Error:

```
 => ERROR [builder 3/7] COPY go.mod ./
 => ERROR [builder 4/7] COPY go.sum ./
```

### ✅ Cause:

These files (`go.mod` and `go.sum`) are not present in the same directory as your Dockerfile when you run `docker build`.

### ✅ Solution:

Before running `docker build`, make sure your project structure looks like this:

```
project-folder/
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
```

If you haven't created the Go module yet, run this:

```bash
go mod init your-module-name
go mod tidy
```

Then build your Docker image **from the same folder**:

```bash
docker build -t go-webapp .
```

---

## 📄 Dockerfile Explanation

```Dockerfile
FROM golang:1.21 AS builder
```
🇺🇸 Build stage using a Go environment.  
🇧🇷 Etapa de build usando ambiente com Go.

```Dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o webserver
```
🇺🇸 Compiles the Go app for Linux with no CGO dependency.  
🇧🇷 Compila a aplicação para Linux sem dependência de CGO.

```Dockerfile
FROM alpine:latest
```
🇺🇸 Runtime stage with a minimal image.  
🇧🇷 Etapa de execução com imagem mínima.

```Dockerfile
EXPOSE 8080
```
🇺🇸 Exposes port 8080 used by `net/http`.  
🇧🇷 Expõe a porta 8080 usada pelo `net/http`.

---

## 🚀 Build and Run the Container

```bash
docker build -t go-webapp .
docker run -d -p 8080:8080 go-webapp
```

Access in your browser:
```
http://localhost:8080/
```

---

## 🔁 Integration with Previous Lecture

This Dockerfile runs the server from **Lecture 01** automatically. As you move forward in the course, this Docker setup will evolve alongside your project.

Este Dockerfile executa o servidor criado na **Lecture 01** automaticamente. À medida que você avança no curso, esse ambiente Docker será atualizado junto com o seu projeto.

---

Happy coding! | Bons estudos!
