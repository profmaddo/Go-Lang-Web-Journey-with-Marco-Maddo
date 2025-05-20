# LECTURE 03 – Static Files and Project Structure

## 🧭 Introduction

🇺🇸 In this lecture, you'll learn how to serve static files (CSS, images, JavaScript) with Go using `http.FileServer`. We'll also introduce the concept of organizing a web project into folders that separate logic (Go code), templates (HTML files), and static assets (CSS, JS, images). This is a key step toward building clean, scalable web applications.

🇧🇷 Nesta aula, você aprenderá a servir arquivos estáticos (CSS, imagens, JavaScript) com Go usando `http.FileServer`. Também vamos apresentar o conceito de organização de projetos web, separando a lógica (código Go), os templates (arquivos HTML) e os assets estáticos (CSS, JS, imagens). Este é um passo fundamental para construir aplicações web limpas e escaláveis.

---

## 📁 Folder Structure | Estrutura de Pastas

```
/seu-projeto/
├── main.go
├── static/
│   └── styles.css
├── templates/
    ├── index.html
    └── about.html
```

🇺🇸 This structure follows the MVC principle (Model-View-Controller) for clarity.  
🇧🇷 Essa estrutura segue o princípio MVC (Modelo-Visão-Controlador) para maior clareza.

---

## 🌐 Static Files Setup

🇺🇸 The Go code uses the following lines to serve static files:

🇧🇷 O código Go utiliza as seguintes linhas para servir os arquivos estáticos:

```go
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

---

## 🧪 How to Run | Como Executar

1. Certifique-se de que o CSS esteja na pasta `static/` e os templates em `templates/`.
2. Execute o servidor:

```bash
go run main.go
```

3. Acesse:

```
http://localhost:8080/        → Página inicial com lista de celulares
http://localhost:8080/about   → Página sobre
```

4. Para ver o CSS diretamente:

```
http://localhost:8080/static/styles.css
```

---

## 🔁 Integration with Previous Lectures

🇺🇸 Builds upon Lecture 02 by styling the rendered templates.  
🇧🇷 Expande a Lecture 02 aplicando estilos visuais nos templates HTML.

🇺🇸 Reinforces the concept of separating concerns (logic, view, static content).  
🇧🇷 Reforça o conceito de separação de responsabilidades (lógica, visual, conteúdo estático).

---

Happy coding! | Bons estudos!
