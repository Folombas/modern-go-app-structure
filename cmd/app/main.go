package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Folombas/modern-go-app-structure/internal/config"
	"github.com/Folombas/modern-go-app-structure/internal/repository"
	"github.com/Folombas/modern-go-app-structure/internal/service"
	"github.com/Folombas/modern-go-app-structure/internal/usecase"
	"github.com/Folombas/modern-go-app-structure/pkg/logger"

	// Явный импорт domain
	_ "github.com/Folombas/modern-go-app-structure/internal/domain"
)

func main() {
	logger.Info("🚀 Starting delivery application")

	// Загрузка конфигурации
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Подключение к базе данных
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	repo, err := repository.NewPostgresRepository(connStr)
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer repo.Close()

	// Инициализация сервисов
	paymentService := service.NewPaymentService()
	orderUsecase := usecase.NewOrderUsecase(repo, paymentService)

	// Создаем пользователя
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	user, err := repo.CreateUser(ctx, "Courier Alex")
	if err != nil {
		logger.Error("❌ Failed to create user: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("👤 User created: %s (ID: %s)\n", user.Name, user.ID)

	// Создаем и обрабатываем заказ
	order, err := orderUsecase.CreateOrder(ctx, user.ID, 690)
	if err != nil {
		logger.Error("❌ Order creation failed: " + err.Error())
		os.Exit(1)
	}
	
	fmt.Printf("✅ Order #%s completed! Amount: %d RUB\n", order.ID, order.Amount)
	logger.Info("🛑 Application stopped gracefully")
}