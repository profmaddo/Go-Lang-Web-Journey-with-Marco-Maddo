# LECTURE 04 – Form Handling and Validation

## 🧭 Introduction

🇺🇸 In this lecture, we focus on handling HTML form submissions and validating input data using Go. You'll learn how to capture data from `POST` requests with `r.FormValue()`, validate required fields, and return appropriate error messages to the user. The final result is a working form that provides feedback and dynamically responds to input.

🇧🇷 Nesta aula, o foco é o tratamento de formulários HTML e validação de dados usando Go. Você aprenderá como capturar dados de requisições `POST` com `r.FormValue()`, validar campos obrigatórios e exibir mensagens de erro apropriadas ao usuário. O resultado final é um formulário funcional com resposta dinâmica e validações simples.

---

## 📄 What You’ll Learn | O que você vai aprender

- ✅ How to use `r.FormValue()` to capture form fields  
- ✅ How to validate required fields  
- ✅ How to return dynamic error messages  
- ✅ How to organize templates for forms and success feedback  
- ✅ Integration of HTML + CSS + logic

---

## 📁 Project Structure

```
/seu-projeto/
├── main.go
├── static/
│   └── styles.css
├── templates/
    ├── form.html
    └── success.html
```

---

## 🌐 Run and Test

```bash
go run main.go
```

Access the form in your browser:

```
http://localhost:8080/register
```

---

## 🧪 What Happens

🇺🇸 When the user submits the form:
- Go reads the input using `r.FormValue()`
- Validates that no field is empty
- Displays errors (if any), or success page

🇧🇷 Quando o usuário envia o formulário:
- Go lê os dados com `r.FormValue()`
- Valida que todos os campos estejam preenchidos
- Exibe erros (se houver) ou uma página de sucesso

---

## 🔁 Integration with Previous Lectures

- 📚 Uses the templating system from Lecture 02  
- 🖼 Styled with `static/styles.css` from Lecture 03  
- 📂 Maintains organized project layout (MVC-inspired)

---

Happy coding! | Bons estudos!
