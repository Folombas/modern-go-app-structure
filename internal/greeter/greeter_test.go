package greeter_test

import (
	"testing"
	"github.com/Folombas/modern-go-app-structure/internal/greeter"
)

func TestSayHello(t *testing.T) {
	expected := "Привет, Тест! Добро пожаловать в мир Go-программирования!"
	result := greeter.SayHello("Тест")
	
	if result != expected {
		t.Errorf("Ожидалось: %s, Получено: %s", expected, result)
	}
}
