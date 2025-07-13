package api

import (
    "fmt"
    "net/http"
    "modern-go-app-structure/internal/api/useragent"
)

func Run() error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Получаем User-Agent из запроса
        userAgent := r.UserAgent()
        
        // Получаем информацию о клиенте
        clientInfo := useragent.GetClientInfo(userAgent)
        
        // Формируем ответ
        message := fmt.Sprintf(`
            <h1>Ваши данные:</h1>
            <p>Операционная система: %s</p>
            <p>Браузер: %s</p>
            <p>Версия браузера: %s</p>
        `, clientInfo.OS, clientInfo.Browser, clientInfo.Version)
        
        // Выводим результат
        w.Write([]byte(message))
    })
    
    return http.ListenAndServe(":8080", nil)
}