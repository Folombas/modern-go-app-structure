package main

import (
	"log"

	"github.com/Folombas/modern-go-app-structure/internal/api"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatalf("Error running API: %v", err)
	}
}

