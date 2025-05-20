#!/bin/bash

echo "🔍 Checking Docker installation..."
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed or not available in PATH."
    exit 1
fi

echo "✅ Docker is installed."

echo ""
echo "📦 Listing Docker containers on port 8080..."
docker ps --filter "publish=8080"

echo ""
echo "🔍 Checking if container is running on port 8080..."
RUNNING=$(docker ps --filter "publish=8080" --format '{{.ID}}')
if [ -z "$RUNNING" ]; then
    echo "❌ No container is currently running on port 8080."
else
    echo "✅ Container is running on port 8080: $RUNNING"
    echo ""
    echo "📄 Logs from container:"
    docker logs "$RUNNING"
fi

echo ""
echo "🔍 Checking if port 8080 is listening on localhost..."
if lsof -i :8080 | grep LISTEN > /dev/null; then
    echo "✅ Port 8080 is active on localhost."
else
    echo "❌ Port 8080 is not active on localhost. The container may have crashed or exited."
fi

echo ""
echo "💡 Tip: Run the container using:"
echo "docker run -d -p 8080:8080 go-webapp"
