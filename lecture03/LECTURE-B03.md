# LECTURE B03 – Dockerfile for Lecture 03 (Static Files and Project Structure)

## 🧭 Introduction

🇺🇸 In this lecture, we updated the Dockerfile to support not only Go templates, but also static files like CSS, images, and JavaScript. Serving static files in Docker requires copying the `/static` folder into the container. We also ensure that the app runs in a clean and efficient Alpine-based environment.

🇧🇷 Nesta aula, atualizamos o Dockerfile para suportar não apenas os templates HTML em Go, mas também arquivos estáticos como CSS, imagens e JavaScript. Para servir esses arquivos no Docker, é necessário copiar a pasta `/static` para dentro do container. Também garantimos que o app seja executado em um ambiente limpo e eficiente baseado em Alpine Linux.

---

## 📄 Dockerfile Explained | Dockerfile Explicado

```dockerfile
FROM golang:1.21 AS builder
```
🇺🇸 Uses official Go image for compilation.  
🇧🇷 Usa imagem oficial do Go para compilar.

```dockerfile
WORKDIR /app
```
🇺🇸 Sets build context path.  
🇧🇷 Define o caminho de trabalho no container de build.

```dockerfile
COPY . .
```
🇺🇸 Copies the entire project.  
🇧🇷 Copia todo o projeto.

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver
```
🇺🇸 Builds a static binary for Alpine.  
🇧🇷 Compila um binário estático compatível com Alpine.

```dockerfile
FROM alpine:latest
```
🇺🇸 Starts a minimal runtime image.  
🇧🇷 Inicia imagem mínima para execução.

```dockerfile
WORKDIR /root/
```
🇺🇸 Sets working directory inside runtime container.  
🇧🇷 Define diretório de trabalho no container final.

```dockerfile
COPY --from=builder /app/webserver .
```
🇺🇸 Copies the Go binary.  
🇧🇷 Copia o binário Go compilado.

```dockerfile
COPY --from=builder /app/templates ./templates
```
🇺🇸 Copies HTML templates.  
🇧🇷 Copia os templates HTML.

```dockerfile
COPY --from=builder /app/static ./static
```
🇺🇸 Copies static files (CSS, JS, images).  
🇧🇷 Copia os arquivos estáticos (CSS, JS, imagens).

```dockerfile
EXPOSE 8080
```
🇺🇸 Informs Docker to expose port 8080.  
🇧🇷 Informa ao Docker que a porta 8080 será usada.

```dockerfile
CMD ["./webserver"]
```
🇺🇸 Defines entrypoint to run the Go web app.  
🇧🇷 Define o comando de inicialização do app Go.

---

## 🚀 Build and Run

```bash
docker build -t go-webapp .
docker run -d -p 8080:8080 go-webapp
```

---

## 🧪 Check with Browser

🇺🇸 After running the container, open:

```
http://localhost:8080/         → Home page with styled list
http://localhost:8080/about    → About page
```

🇧🇷 Após executar o container, acesse:

```
http://localhost:8080/         → Página inicial com lista estilizada
http://localhost:8080/about    → Página sobre
```

---

Happy coding! | Bons estudos!
