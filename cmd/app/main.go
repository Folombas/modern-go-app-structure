package main

import (
	"fmt"
	"log"
	"github.com/Folombas/modern-go-app-structure/internal/greeter"
)

func main() {
	fmt.Println("🚀 Приложение успешно запущено!")
	log.Println("🔧 Логирование работает корректно")
	
	// Используем наш пакет для приветствия
	message := greeter.SayHello("Гоша")
	fmt.Println(message)
	
	fmt.Println("✨ Попробуй изменить код и запустить снова: go run ./cmd/app")
}
