# LECTURE 06 – User Registration and Password Hashing

## 🧭 Introduction

🇺🇸 In this lecture, you'll learn how to build a secure user registration form in Go using **bcrypt** to hash passwords before storing them in a SQLite database. This introduces a critical web development topic: **security**. Plain text passwords are never stored — only hashes.

🇧🇷 Nesta aula, você aprenderá como construir um formulário seguro de cadastro de usuários em Go, utilizando **bcrypt** para criptografar senhas antes de salvá-las no banco de dados SQLite. Isso introduz um tema fundamental no desenvolvimento web: **segurança**. Senhas em texto puro nunca são armazenadas — apenas os hashes.

---

## 🧰 Skills Gained

- ✅ Create a user form with validation  
- ✅ Use `bcrypt.GenerateFromPassword` to hash  
- ✅ Use SQLite to store users securely  
- ✅ Check for duplicate emails  
- ✅ Display clear messages to users

---

## 🔐 About bcrypt

🇺🇸 `bcrypt` is a password hashing function designed for security. It's slow and resistant to brute force attacks. In Go, we use the package:

```go
import "golang.org/x/crypto/bcrypt"
```

🇧🇷 `bcrypt` é uma função de hashing de senhas feita para segurança. É lenta e resistente a ataques de força bruta. Em Go, usamos o pacote:

```go
import "golang.org/x/crypto/bcrypt"
```

---

## 🌐 How to Test

```bash
go run main.go
```

Then open:

```
http://localhost:8080/user/register
```

---

## 📁 Project Structure

```
/seu-projeto/
├── main.go
├── /templates/
│   ├── register.html
│   └── success.html
├── /static/styles.css
├── /database/users.db (volume externo)
```

---

Happy coding! | Bons estudos!
