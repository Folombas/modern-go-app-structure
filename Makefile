# Простейший Makefile для Go-проекта

run: # Запуск приложения
	go run cmd/main.go

test: # Запуск тестов
	go test ./...

build: # Сборка бинарника
	go build -o app-bin cmd/main.go

clean: # Удаление бинарника
	rm -f app-bin

lint: # Проверка кода через golangci-lint
	golangci-lint run