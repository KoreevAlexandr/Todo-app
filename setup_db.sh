#!/bin/bash

# Checking PostgreSQL install
if ! command -v psql &> /dev/null; then
    echo "PostgreSQL is not installed:"
    echo "sudo pacman -S postgresql"
    exit 1
fi

# Checking if PostgreSQL is running
if ! systemctl is-active --quiet postgresql; then
    sudo systemctl start postgresql
    sudo systemctl enable postgresql
fi

# Creating User Database
echo "User DB creation..."
sudo -u postgres psql -c "CREATE DATABASE todo_app;"
sudo -u postgres psql -c "CREATE USER postgres WITH PASSWORD 'password';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE todo_app TO postgres;"

echo "go run cmd/main.go"
