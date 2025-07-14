package api

import (
	"fmt"
	"net/http"

	"github.com/Folombas/modern-go-app-structure/internal/useragent"
)

func Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		clientIP := r.Header.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = r.RemoteAddr // Резервный вариант
		}

		clientInfo := useragent.GetClientInfo(userAgent, clientIP)

		message := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head><title>Client Info</title></head>
        <body>
            <h1>Ваши данные:</h1>
            <p>Тип устройства: %s</p>
            <p>Операционная система: %s</p>
            <p>Браузер: %s %s</p>
            <p>Ядер процессора на сервере: %d</p>
            <p>Тип подключения на сервере: %s</p>
        </body>
        </html>
    `, clientInfo.DeviceType, clientInfo.OS, clientInfo.Browser, clientInfo.Version, clientInfo.Cores, clientInfo.ConnectionType)

		w.Write([]byte(message))
	})

	return http.ListenAndServe(":8080", nil)
}
