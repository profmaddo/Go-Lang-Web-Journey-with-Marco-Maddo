# LECTURE B02 – Dockerfile for Lecture 02 (HTML Templating with Go)

## 🧭 Introduction

🇺🇸 In this Docker setup, we extend the containerization of our Go web application by including a template folder with HTML files. This structure supports dynamic rendering via Go's `html/template` engine. The Dockerfile is built with a multi-stage strategy, producing a lightweight and efficient final image ready to run.

🇧🇷 Neste setup Docker, expandimos a containerização da aplicação web em Go incluindo a pasta de templates com arquivos HTML. Essa estrutura permite renderização dinâmica com o pacote `html/template`. O Dockerfile utiliza uma estratégia multi-stage para gerar uma imagem final leve e eficiente, pronta para execução.

---

## 📄 Dockerfile Explained | Dockerfile Explicado

```dockerfile
FROM golang:1.21 AS builder
```
🇺🇸 Uses the official Go image to compile the application.  
🇧🇷 Usa a imagem oficial do Go para compilar a aplicação.

```dockerfile
WORKDIR /app
```
🇺🇸 Sets the working directory inside the build container.  
🇧🇷 Define o diretório de trabalho dentro do container de build.

```dockerfile
COPY . .
```
🇺🇸 Copies all files from your project into the container.  
🇧🇷 Copia todos os arquivos do projeto para dentro do container.

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver
```
🇺🇸 Builds a static binary compatible with Alpine Linux.  
🇧🇷 Compila o binário de forma estática, compatível com Alpine.

```dockerfile
FROM alpine:latest
```
🇺🇸 Starts a clean minimal image for runtime.  
🇧🇷 Inicia uma imagem limpa e mínima para execução.

```dockerfile
WORKDIR /root/
```
🇺🇸 Sets working directory inside the runtime container.  
🇧🇷 Define o diretório de trabalho no container final.

```dockerfile
COPY --from=builder /app/webserver .
```
🇺🇸 Copies the compiled Go binary from builder stage.  
🇧🇷 Copia o binário compilado da etapa de build.

```dockerfile
COPY --from=builder /app/templates ./templates
```
🇺🇸 Copies the folder with HTML templates into the container.  
🇧🇷 Copia a pasta com os templates HTML para dentro do container.

```dockerfile
EXPOSE 8080
```
🇺🇸 Informs Docker that the container listens on port 8080.  
🇧🇷 Informa ao Docker que o container escuta na porta 8080.

```dockerfile
CMD ["./webserver"]
```
🇺🇸 Sets the command to execute the Go web server when container starts.  
🇧🇷 Define o comando para executar o servidor Go quando o container iniciar.

---

## 🚀 How to Build and Run

```bash
docker build -t go-webapp .
docker run -d -p 8080:8080 go-webapp
```

---

## 🌐 How to Access the Web Server

🇺🇸 After running, visit:  
🇧🇷 Após executar, acesse:

```
http://localhost:8080/
http://localhost:8080/about
```

---

## 🧪 Diagnose with Script

🇺🇸 To confirm if your server is running correctly and listening on port 8080, use:  
🇧🇷 Para confirmar se o servidor está rodando corretamente na porta 8080, use:

```bash
./diagnose.sh
```

Make sure the script has execute permission:  
Certifique-se de que o script tem permissão de execução:

```bash
chmod +x diagnose.sh
```

---

Happy coding! | Bons estudos!
