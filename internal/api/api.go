package api

import (
    "fmt"
    "log"
    "net/http"
)

func Run() error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Теперь проверяем ошибку
        _, err := fmt.Fprintf(w, "Hello from Modern-go-app-structure!")
        if err != nil {
            log.Printf("Error writing response: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    })
    return http.ListenAndServe(":8080", nil)
}