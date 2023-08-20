#!/bin/bash

# This is a simple database migration script for a Go application.
# It runs database migrations using the "migrate" tool.

# Path to the "migrate" binary
MIGRATE_BINARY="./migrate"  # Adjust the path to the "migrate" binary

# Database connection URL
DB_URL="postgres://username:password@localhost:5432/mydb?sslmode=disable"

# Directory containing migration files
MIGRATIONS_DIR="./migrations"

# Run the migrations
$MIGRATE_BINARY -database "$DB_URL" -path "$MIGRATIONS_DIR" up
