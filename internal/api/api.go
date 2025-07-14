package api

import (
	"fmt"
	"net/http"

	"github.com/Folombas/modern-go-app-structure/internal/useragent"
)

func Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    userAgent := r.UserAgent()
    
    // Получаем IP-адрес клиента
    clientIP := r.Header.Get("X-Forwarded-For")
    if clientIP == "" {
        clientIP = r.RemoteAddr // Резервный вариант
    }

    clientInfo := useragent.GetClientInfo(userAgent, clientIP)

    message := fmt.Sprintf(`
        <h1>Ваши данные:</h1>
        <p>Тип устройства: %s</p>
        <p>Операционная система: %s</p>
        <p>Модель устройства: %s</p>
        <p>Браузер: %s %s</p>
        <p>Тип подключения: %s</p>
    `, 
        clientInfo.DeviceType, 
        clientInfo.OS, 
        clientInfo.Model, 
        clientInfo.Browser, 
        clientInfo.Version,
        clientInfo.ConnectionType,
    )

    w.Write([]byte(message))
})

	return http.ListenAndServe(":8080", nil)
}
