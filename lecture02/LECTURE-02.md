# LECTURE 02 – HTML Templating with Go

## 🧭 Introduction

🇺🇸 In this lecture, you will learn how to serve dynamic HTML content using Go's built-in `html/template` package. This is a core skill for any web application that displays data. You'll build a simple system to list used phones by rendering a Go `struct` inside HTML using `{{range}}` and `{{.}}`. By the end of this lecture, you’ll understand how to organize your templates into separate files and pass data from Go to the HTML frontend dynamically.

🇧🇷 Nesta aula, você aprenderá como servir conteúdo HTML dinâmico usando o pacote `html/template` do próprio Go. Esta é uma habilidade essencial para qualquer aplicação web que exibe dados. Vamos construir um sistema simples para listar celulares usados renderizando uma `struct` Go no HTML com `{{range}}` e `{{.}}`. Ao final da aula, você saberá como organizar seus templates em arquivos separados e como passar dados do Go para o frontend HTML dinamicamente.

---

## 📄 What You’ll Learn | O que você vai aprender

- ✅ How to use the `html/template` package  
- ✅ How to pass structs and slices to templates  
- ✅ How to use `{{range}}`, `{{.}}`, and `{{end}}`  
- ✅ How to structure templates with `/templates` folder  

---

## 🧪 Run the Example | Execute o Exemplo

1. Crie uma pasta chamada `templates` no mesmo nível do `main.go`
2. Salve os arquivos `index.html` e `about.html` dentro da pasta `templates`
3. Rode a aplicação:

```bash
go run main.go
```

4. Abra no navegador:

```
http://localhost:8080/        → Home com a lista de celulares
http://localhost:8080/about   → Página sobre o projeto
```

---

## 💡 Template Preview (index.html)

```html
<ul>
    {{range .}}
        <li><strong>{{.Brand}}</strong> - {{.Model}}: R$ {{.Price}}</li>
    {{else}}
        <li>No phones available.</li>
    {{end}}
</ul>
```

🇺🇸 This code loops through a list of phones and displays each one.  
🇧🇷 Esse trecho percorre a lista de celulares e exibe cada item.

---

## 🔁 Integration with Lecture 01

- 🇺🇸 Replaces the plain text response with structured HTML views  
- 🇧🇷 Substitui as respostas em texto simples por páginas HTML estruturadas  
- 🇺🇸 Starts introducing the concept of Model → View → Controller  
- 🇧🇷 Introduz o conceito de Modelo → Visão → Controlador (MVC)

---

Happy coding! | Bons estudos!
