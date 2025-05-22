# LECTURE 08 – Best Practices for Modularization in Go

## 🧭 Introduction

🇺🇸 In this lecture, you will refactor your Go Web project by applying modularization best practices. You’ll create folders for each lecture, separate handlers into packages, and improve the structure of your `main.go` file to serve as a central index. This will help you build scalable, clean, and maintainable applications.

🇧🇷 Nesta aula, você irá refatorar seu projeto Web em Go aplicando boas práticas de modularização. Você criará pastas para cada lecture, separará os handlers em pacotes e melhorará a estrutura do arquivo `main.go`, que passará a servir como índice central. Isso ajudará a construir aplicações escaláveis, limpas e de fácil manutenção.

---

## 🧪 What You’ll Do | O que você vai fazer

- ✅ Refactor the existing project using packages  
- ✅ Organize each lecture in its own folder (e.g., `lecture01/handler.go`)  
- ✅ Use `go mod init golangweb` to initialize the module  
- ✅ Create a modular and clean `main.go`  
- ✅ Maintain clean Git commits and folder structure

---

## 🗂️ Recommended Project Structure | Estrutura Recomendada

```
/golangweb/
├── go.mod
├── main.go
├── /lecture01/
│   └── handler.go
├── /lecture02/
│   └── handler.go
├── /lecture03/
│   └── handler.go
...
├── /lecture07/
│   └── handler.go
├── /templates/
```

---

## 🧰 How to Refactor | Como Refatorar

```bash
# Step 1: Initialize Go module
go mod init golangweb

# Step 2: Create folders for lectures
mkdir lecture01 lecture02 lecture03 lecture04 lecture05 lecture06 lecture07

# Step 3: Move handler files
mv lecture01.go lecture01/handler.go
# Repita o processo para cada lecture

# Step 4: Adjust imports in main.go
import "golangweb/lecture01"

# Step 5: Run the project
go run main.go
```

---

## 🔁 Integration with Previous Lectures

🇺🇸 This modularization includes all previous lectures from 01 to 07, making your project easier to scale as you move forward.

🇧🇷 Esta modularização inclui todas as lectures anteriores de 01 a 07, tornando seu projeto mais fácil de escalar conforme você avança.

---

## 🛠 Git Tips | Dicas de Git

```bash
git init
git add .
git commit -m "Refactor project with modular lecture packages"
git remote add origin https://github.com/profmaddo/Go-Lang-Web-Journey-with-Marco-Maddo.git
git push -u origin main
```

---

## 🙋 Need Help?

🇺🇸 If you run into problems, feel free to [open an issue](https://github.com/profmaddo/Go-Lang-Web-Journey-with-Marco-Maddo.git/issues).

🇧🇷 Se tiver dificuldades, sinta-se à vontade para [abrir uma issue](https://github.com/profmaddo/Go-Lang-Web-Journey-with-Marco-Maddo.git/issues).

---

## 👨‍🏫 About This Course

🇺🇸 This course is a voluntary project maintained by [Professor Marco Maddo](https://www.linkedin.com/in/marcomaddo/) and sponsored by [TSSTI Tecnologia](https://www.linkedin.com/company/tssti/?viewAsMember=true), a company specialized in E-Learning.

🇧🇷 Este curso é um projeto voluntário mantido pelo [Professor Marco Maddo](https://www.linkedin.com/in/marcomaddo/) e patrocinado pela [TSSTI Tecnologia](https://www.linkedin.com/company/tssti/?viewAsMember=true), empresa especializada em E-Learning.

---

👨‍🏫 Happy coding! | Bons estudos! 🚀
