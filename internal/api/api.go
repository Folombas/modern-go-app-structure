// internal/api/api.go
package api

import (
    "fmt"
    "net/http"
)

func Run() error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Modern-go-app-structure!")
    })
    return http.ListenAndServe(":8080", nil)
}