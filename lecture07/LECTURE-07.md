# LECTURE 07 – User Login and Sessions

## 🧭 Introduction

🇺🇸 In this lecture, you'll learn how to build a user login system using Go. It includes form submission, secure password comparison using `bcrypt`, and the use of **cookies** to manage user sessions. You'll also see how to **protect routes** that require authentication using middleware-like functions.

🇧🇷 Nesta aula, você aprenderá a criar um sistema de login de usuários em Go. O projeto inclui envio de formulário, comparação segura de senha com `bcrypt`, e o uso de **cookies** para gerenciar sessões de usuários. Você também verá como **proteger rotas** que exigem autenticação usando funções similares a middlewares.

---

## 🔐 What You’ll Learn | O que você vai aprender

- ✅ How to validate user credentials with bcrypt  
- ✅ How to create and read secure cookies  
- ✅ How to implement a protected dashboard  
- ✅ How to destroy sessions (logout)  
- ✅ How to integrate login with SQLite user database

---

## 🌐 How it works

1. User accesses `/login` and submits email + password
2. If credentials are valid, a cookie `session_email` is set
3. Access to `/dashboard` checks if the cookie exists
4. If not logged in, user is redirected to `/login`
5. Clicking **Logout** deletes the cookie and redirects

---

## 🧪 How to Test

```bash
go run main.go
```

Then visit:

```
http://localhost:8080/login
```

Try logging in with a user that you created in Lecture 6.

---

## 📁 Project Structure

```
/seu-projeto/
├── main.go
├── /templates/
│   ├── login.html
│   ├── dashboard.html
│   └── error.html
├── /database/users.db (mounted volume)
```

---

## 🔁 Integration with Previous Lectures

- 🧠 Lecture 06: Reuses the user data from registration
- 📄 Lecture 05: SQLite database continues to be used
- 🎨 Lecture 03: CSS is loaded from `/static/`
- 🧱 Lecture 02: Uses HTML templates with `{{.}}` data binding

---

Happy coding! | Bons estudos!
