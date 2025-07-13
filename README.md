# Modern-go-app-structure

Это учебное приложение, демонстрирующее современную структуру Go-проекта.

## Описание

Приложение включает:

- HTTP-сервер (в папке `internal/api`)
- Бизнес-логику (в папке `internal/service`)
- Тесты (в папке `tests`)
- Конфигурации и утилиты (в папке `pkg`)

## Установка

```bash
go mod tidy
```

## Пример вывода
```bash
$ go run cmd/main.go
Server started on port 8080