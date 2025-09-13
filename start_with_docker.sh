#!/bin/bash

# Checking Docker installation
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed:"
    echo "sudo pacman -S docker docker-compose"
    exit 1
fi

# Checking docker-compose installation
if ! command -v docker-compose &> /dev/null; then
    echo "Docker Compose is not installed:"
    echo "sudo pacman -S docker-compose"
    exit 1
fi

# DB init
./init_db.sh

# Go reqs
go mod tidy

# run app
go run cmd/main.go
