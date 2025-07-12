package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/Folombas/modern-go-app-structure/internal/greeter"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	startTime      = time.Now()
	clients        = make(map[*websocket.Conn]bool)
	clientsMutex   sync.Mutex
	upgrader       = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Разрешаем все источники (для разработки)
		},
	}
)

func init() {
	log.Println("🛠 Инициализация приложения...")
}

func main() {
	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using defaults")
	}

	// Конфигурация
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	if port == 0 {
		port = 8080
	}
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Инициализация логгера
	log.SetPrefix("[APP] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("🚀 Приложение инициализируется в окружении '%s' на порту %d...", env, port)

	// Использование пакета greeter
	message, err := greeter.SayHello("Гоша")
	if err != nil {
		log.Fatalf("Ошибка в пакете greeter: %v", err)
	}
	fmt.Println(message)

	// Настройка HTTP-сервера
	mux := http.NewServeMux()
	
	// Веб-страница
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<title>Modern Go App</title>
	<meta charset="UTF-8">
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
		#uptime { 
			font-weight: bold;
			color: #2c3e50;
			font-size: 1.2em;
		}
	</style>
	<script>
		const ws = new WebSocket("ws://" + window.location.host + "/ws");
		
		ws.onmessage = function(event) {
			const data = JSON.parse(event.data);
			if (data.uptime) {
				document.getElementById('uptime').textContent = data.uptime;
			}
		};
		
		ws.onerror = function(error) {
			console.error("WebSocket Error:", error);
		};
	</script>
</head>
<body>
	<h1>Привет от Go! 🚀</h1>
	<div class="info">
		<p>Сервер работает: <span id="uptime">%s</span></p>
		<p>Окружение: <span class="env-badge">%s</span></p>
		<p>Порт: <strong>%d</strong></p>
		<p>Версия: <strong>%s</strong></p>
		<p><small>Время обновляется в реальном времени с сервера</small></p>
	</div>
</body>
</html>
		`, getEnvColor(env), formatDuration(time.Since(startTime)), env, port, getVersion())
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	})

	// Health-check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"status": "ok", 
			"timestamp": "%s",
			"version": "%s",
			"uptime": "%s"
		}`, time.Now().Format(time.RFC3339), getVersion(), formatDuration(time.Since(startTime)))
	})

	// WebSocket endpoint
	mux.HandleFunc("/ws", handleWebSocket)

	// Запуск сервера
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запуск рассылки обновлений времени
	go broadcastUptime()

	go func() {
		log.Printf("🌐 HTTP-сервер запущен на порту %d", port)
		log.Printf("👉 Откройте в браузере: http://localhost:%d", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	// Обработка сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("🛑 Получен сигнал завершения...")
	
	// Корректное завершение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Принудительное завершение: %v", err)
	}
	
	log.Println("✅ Сервер корректно остановлен")
}

// Обработчик WebSocket соединений
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	// Регистрируем нового клиента
	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	// Отправляем начальное значение
	conn.WriteJSON(map[string]string{
		"uptime": formatDuration(time.Since(startTime)),
	})

	// Ожидаем сообщений (клиент может отправить запрос на обновление)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
	
	// Удаляем клиента при отключении
	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()
}

// Рассылка обновлений всем подключенным клиентам
func broadcastUptime() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		uptime := formatDuration(time.Since(startTime))
		
		clientsMutex.Lock()
		for client := range clients {
			err := client.WriteJSON(map[string]string{
				"uptime": uptime,
			})
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMutex.Unlock()
	}
}

// Форматирование времени в ЧЧ:ММ:СС
func formatDuration(d time.Duration) string {
	total := int(d.Seconds())
	hours := total / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// Вспомогательные функции (getEnvColor, getVersion остаются без изменений)