# LECTURE B01 – Docker Setup for Go Web App

## 🧭 Introduction

In this optional but essential lecture, we’ll configure Docker to build and run our Go web application. Docker creates consistent development and deployment environments across systems. This Docker setup supports all lectures from Lecture 01 onward, and will be continuously updated as the course progresses.

Nesta aula opcional, porém essencial, vamos configurar o Docker para compilar e executar nossa aplicação web escrita em Go. O Docker cria ambientes de desenvolvimento e deploy consistentes em qualquer sistema operacional. Essa configuração do Docker será mantida ao longo das próximas aulas do curso.

---

## 🛠 How to Build and Run the Application

### Step 1 – Build the Docker image

```bash
docker build -t go-webapp .
```

### Step 2 – Run the container on port 8080

```bash
docker run -d --name go-webapp-container -p 8080:8080 go-webapp
```

---

## 🌐 Accessing the Web Server

After the container is running, open your browser and navigate to:

```
http://localhost:8080/
```

If you created an `index.html` and updated `main.go` to serve it, it will load automatically.

You can also test using:

```bash
curl http://localhost:8080/
```

---

## 🧩 Scripts Included

### 📄 `build-and-run.sh`

This script automates the image build and container launch:

```bash
chmod +x build-and-run.sh
./build-and-run.sh
```

What it does:
- Builds the Docker image
- Removes any previously running container
- Starts the new container and publishes it on port 8080

---

### 📄 `diagnose.sh`

This script helps diagnose if the Docker container is running correctly:

```bash
chmod +x diagnose.sh
./diagnose.sh
```

What it checks:
- Docker installation
- Containers using port 8080
- Port 8080 status
- Logs from the running container

Use it when:
- The container doesn’t appear to work
- You see errors like `connection refused` or nothing loads in the browser

---

## 🧯 Troubleshooting

### ❌ COPY go.mod / go.sum fails

**Cause**: Files not present  
**Fix**: Run `go mod init` and ensure `go.mod` exists before build. Remove `COPY go.sum` if unnecessary.

---

### ❌ RUN go mod download fails

**Cause**: No dependencies  
**Fix**: Remove this line from Dockerfile if `go.sum` doesn't exist

---

### ❌ Go version mismatch

**Error**:
```
go: go.mod requires go >= 1.23.0 (running go 1.21.13)
```

**Fix**: Edit `go.mod` to require Go 1.21:
```go
go 1.21
```

---

### ❌ exec ./webserver: no such file or directory

**Cause**: Alpine can't run dynamically linked binaries  
**Fix**: Build with:

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver
```

---

### ❌ ERR_CONNECTION_REFUSED in browser

**Cause**: Container not running or port not mapped  
**Fix**: Make sure to run:

```bash
docker run -d -p 8080:8080 go-webapp
```

Check with:

```bash
docker ps
docker logs <container_id>
```

---

Happy coding! | Bons estudos!
