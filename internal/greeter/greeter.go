package greeter

import (
	"errors"
	"strings"
)

// SayHello возвращает персонализированное приветствие
// Если имя пустое или содержит только пробелы - возвращает ошибку
func SayHello(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("имя не может быть пустым")
	}
	return "Привет, " + trimmed + "! Добро пожаловать в мир Go-программирования!", nil
}

// generateWelcome генерирует базовое приветствие (внутренняя функция)
func generateWelcome() string {
	return "Добро пожаловать в мир Go-программирования!"
}