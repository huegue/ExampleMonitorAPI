#!/bin/bash

DB_FILE="./data/products.db"
BINARY="./monitor_api"

if [ ! -f "$DB_FILE" ]; then
    echo "Database file not found. Creating database..."
    "$BINARY" --createdb
fi

echo "Starting the application..."
"$BINARY" --start



