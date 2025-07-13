package api

import (
	"fmt"
	"net/http"

	"github.com/Folombas/modern-go-app-structure/internal/useragent"
)

func Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		clientInfo := useragent.GetClientInfo(userAgent)

		message := fmt.Sprintf(`
            <!DOCTYPE html>
            <html>
            <head><title>Client Info</title></head>
            <body>
                <h1>Ваши данные:</h1>
                <p>Тип вашего устройства: %s</p>
                <p>Ваша операционная система: %s</p>
                <p>Количество ядер процессора: %d</p>
                <p>Модель вашего устройства: %s</p>
                <p>Ваш Web-браузер: %s</p>
                <p>Версия вашего web-браузера: %s</p>
            </body>
            </html>
        `, clientInfo.DeviceType, clientInfo.OS, clientInfo.Cores, clientInfo.Model, clientInfo.Browser, clientInfo.Version)

		w.Write([]byte(message))
	})

	return http.ListenAndServe(":8080", nil)
}
