# GIT-COMMANDS.md

## 🇺🇸 Basic Git Commands for Daily Use  
## 🇧🇷 Comandos Git Básicos para Uso Diário

---

### ✅ Clone the Repository | Clonar o Repositório

```bash
git clone https://github.com/profmaddo/Go-Lang-Web-Journey-with-Marco-Maddo.git
cd Go-Lang-Web-Journey-with-Marco-Maddo
```

---

### ✅ Check Status | Verificar o Status

```bash
git status
```

---

### ✅ List Commits | Listar os Commits

```bash
git log --oneline --graph --all
```

---

### ✅ Make a Commit | Efetuar um Commit

```bash
git add .
git commit -m "Descrição clara do que foi alterado"
```

---

### ✅ Push to Remote | Efetuar um Push

```bash
git push origin main
# Ou para branch atual:
git push origin nome-da-sua-branch
```

---

### ✅ Pull Updates | Efetuar um Pull

```bash
git pull origin main
# Ou da sua branch
git pull origin nome-da-sua-branch
```

---

### ✅ Update Local Project | Efetuar Atualização do Projeto Local

```bash
git fetch --all
git pull
```

---

### ✅ Merge Branch | Efetuar Merge

```bash
git checkout main
git merge nome-da-sua-branch
```

---

### ✅ Create Pull Request | Criar um Pull Request

1. Push your branch to GitHub:
```bash
git push origin nome-da-sua-branch
```

2. Go to your repository on GitHub  
3. You’ll see a message offering to create a Pull Request  
4. Click "Compare & pull request"  
5. Add title and description  
6. Click "Create pull request"

---

Happy coding! | Bons estudos! 🚀
