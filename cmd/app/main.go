package main

import (
	"fmt"
	"github.com/Folombas/modern-go-app-structure/internal/domain"
	"github.com/Folombas/modern-go-app-structure/internal/service"
	"github.com/Folombas/modern-go-app-structure/internal/usecase"
	"github.com/Folombas/modern-go-app-structure/pkg/logger"
)

func main() {
	logger.Info("🚀 Starting delivery application")
	
	// Инициализация зависимостей
	userRepo := domain.NewUserRepository()
	orderRepo := domain.NewOrderRepository()
	paymentService := service.NewPaymentService()
	
	// Сценарий использования
	orderUsecase := usecase.NewOrderUsecase(orderRepo, paymentService)
	
	// Создаем пользователя
	user := userRepo.CreateUser("Courier Alex")
	fmt.Printf("👤 User created: %s\n", user.Name)
	
	// Создаем и обрабатываем заказ
	order, err := orderUsecase.CreateOrder(user.ID, 690)
	if err != nil {
		logger.Error("❌ Order creation failed: " + err.Error())
		return
	}
	
	fmt.Printf("✅ Order #%s completed! Amount: %d RUB\n", order.ID, order.Amount)
}