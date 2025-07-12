package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Folombas/modern-go-app-structure/internal/greeter"
	"github.com/joho/godotenv"
)

var startTime = time.Now()

func init() {
	log.Println("🛠 Инициализация приложения...")
}

func main() {
	// 0. Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using defaults")
	}

	// 1. Обработка конфигурации
	// Получаем порт из переменных окружения или используем значение по умолчанию
	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		portStr = "8080" // Значение по умолчанию
	}

	// Преобразуем строку в число
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Неверный формат порта: %v", err)
	}

	// Получаем окружение
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Значение по умолчанию
	}

	// 2. Инициализация логгера
	log.SetPrefix("[APP] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("🚀 Приложение инициализируется в окружении '%s' на порту %d...", env, port)

	// 3. Использование нашего пакета greeter
	message, err := greeter.SayHello("Гоша")
	if err != nil {
		log.Fatalf("Ошибка в пакете greeter: %v", err)
	}
	fmt.Println(message)

	// 4. Настройка HTTP-сервера
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>Modern Go App</title>
	<style>
		body { font-family: Arial, sans-serif; text-align: center; padding: 50px; }
		h1 { color: #2c3e50; }
		.info { background: #f8f9fa; padding: 20px; border-radius: 10px; display: inline-block; }
		.env-badge { 
			background-color: %s; 
			color: white; 
			padding: 2px 10px; 
			border-radius: 10px; 
			font-size: 0.8em;
		}
	</style>
</head>
<body>
	<h1>Привет от Go! 🚀</h1>
	<div class="info">
		<p>Сервер работает: <strong>%s</strong></p>
		<p>Окружение: <span class="env-badge">%s</span></p>
		<p>Порт: <strong>%d</strong></p>
		<p>Версия: <strong>%s</strong></p>
	</div>
</body>
</html>
		`, getEnvColor(env), time.Since(startTime).Round(time.Second), env, port, getVersion())
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "ok", 
			"timestamp": "%s",
			"version": "%s",
			"uptime": "%s"
		}`, time.Now().Format(time.RFC3339), getVersion(), time.Since(startTime).Round(time.Second))
	})

	// 5. Запуск сервера в отдельной горутине
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("🌐 HTTP-сервер запущен на порту %d", port)
		log.Printf("👉 Откройте в браузере: http://localhost:%d", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	// 6. Обработка сигналов для корректного завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("🛑 Получен сигнал завершения...")
	
	// Создаем контекст с таймаутом для завершения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Принудительное завершение: %v", err)
	}
	
	log.Println("✅ Сервер корректно остановлен")
}

// Вспомогательная функция для цвета окружения
func getEnvColor(env string) string {
	switch env {
	case "production":
		return "#e74c3c" // Красный
	case "staging":
		return "#f39c12" // Оранжевый
	default:
		return "#2ecc71" // Зеленый
	}
}

// Вспомогательная функция для получения версии приложения
func getVersion() string {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		return "1.0.0-dev"
	}
	return version
}