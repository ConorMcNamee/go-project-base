#!/bin/bash

echo "Starting Docker Compose service..."
docker-compose --env-file .env up -d

# Check if Docker Compose starts successfully
if [ $? -eq 0 ]; then
    echo "Docker Compose started successfully."
else
    echo "Failed to start Docker Compose."
    exit 1
fi

echo "Running Backend..."

go run .

echo "Application Closing"

docker-compose down