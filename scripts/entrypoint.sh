#!/bin/sh
echo "🔄 Применяем миграции БД..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f /app/migrations/001_init_schema.up.sql
echo "🚀 Запускаем приложение..."
./delivery-app