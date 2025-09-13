#!/bin/bash

# Docker init script

# Checking if PostgreSQL is running
if ! docker-compose ps postgres | grep -q "Up"; then
    docker-compose up -d postgres
    sleep 10
fi

# Checking if PostgreSQL is ready
until docker-compose exec postgres pg_isready -U postgres -d todo_app; do
    echo "PostgreSQL isn't ready..."
    sleep 2
done

# Checking if Tables exists
TABLE_COUNT=$(docker-compose exec postgres psql -U postgres -d todo_app -t -c "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public';" | tr -d ' ')

if [ "$TABLE_COUNT" -eq "0" ]; then
    echo "Creating tables..."
    docker-compose exec postgres psql -U postgres -d todo_app -f /docker-entrypoint-initdb.d/000001_init.up.sql
    echo "Tables created!"
else
    echo "Tables exists($TABLE_COUNT tables)"
fi

echo "Database is ready!"
