#!/bin/bash

# Database connection details
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="root"
DB_USER="root"
DB_PASSWORD="secret"

# Path to your SQL script
SQL_SCRIPT="users.sql"

# Export the password to avoid being prompted
export PGPASSWORD=$DB_PASSWORD

# Run the SQL script
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f $SQL_SCRIPT

# Unset the password after use for security reasons
unset PGPASSWORD
