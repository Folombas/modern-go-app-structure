#!/bin/sh

set -e

echo "🔄 Applying database migrations..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f /app/migrations/001_init_schema.up.sql

echo "🚀 Starting application..."
exec /app/delivery-app