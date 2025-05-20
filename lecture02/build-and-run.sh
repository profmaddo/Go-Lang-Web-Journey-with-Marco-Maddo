#!/bin/bash

echo "🔧 Building Docker image 'go-webapp'..."
docker build -t go-webapp .

if [ $? -ne 0 ]; then
    echo "❌ Build failed. Check Dockerfile and source files."
    exit 1
fi

echo ""
echo "🧼 Removing existing container (if any)..."
docker rm -f go-webapp-container &> /dev/null

echo ""
echo "🚀 Running container on port 8080..."
docker run -d --name go-webapp-container -p 8080:8080 go-webapp

echo ""
echo "✅ Application is running at: http://localhost:8080"
