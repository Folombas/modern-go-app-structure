package usecase

import (
	"context"
	"github.com/Folombas/modern-go-app-structure/internal/domain"
	"github.com/Folombas/modern-go-app-structure/internal/repository"
	"github.com/Folombas/modern-go-app-structure/internal/service"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, userID string, amount int) (*domain.Order, error)
}

type orderUsecase struct {
	orderRepo   repository.OrderRepository // Используем интерфейс
	paymentServ service.PaymentService
}

func NewOrderUsecase(
	orderRepo repository.OrderRepository, // Используем интерфейс
	paymentServ service.PaymentService,
) OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		paymentServ: paymentServ,
	}
}

func (uc *orderUsecase) CreateOrder(ctx context.Context, userID string, amount int) (*domain.Order, error) {
	order, err := uc.orderRepo.CreateOrder(ctx, userID, amount)
	if err != nil {
		return nil, err
	}
	
	if err := uc.paymentServ.ProcessPayment(amount); err != nil {
		return nil, err
	}
	
	return order, nil
}

func (uc *orderUsecase) GetRecentOrders(ctx context.Context, limit int) ([]*domain.Order, error) {
	// В реальном приложении здесь была бы бизнес-логика
	// Например, проверка прав доступа
	return uc.orderRepo.GetRecentOrders(ctx, limit)
}