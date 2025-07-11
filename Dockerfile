# Билд-стадия
FROM golang:1.22-alpine AS builder

# Установка зависимостей для сборки
RUN apk add --no-cache git ca-certificates build-base

# Рабочая директория
WORKDIR /app

# Копируем файлы модулей и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/delivery-app ./cmd/app

# Финальная стадия
FROM alpine:latest

# Устанавливаем инструменты для миграций
RUN apk add --no-cache postgresql-client

# Копируем бинарник из стадии сборки
COPY --from=builder /app/bin/delivery-app /app/delivery-app
# Копируем конфиги
COPY configs /app/configs
# Копируем миграции
COPY migrations /app/migrations
# Копируем скрипты
COPY scripts /app/scripts

# Рабочая директория
WORKDIR /app

# Даем права на выполнение скриптов
RUN chmod +x /app/scripts/*.sh

# Команда запуска
CMD ["/app/scripts/entrypoint.sh"]