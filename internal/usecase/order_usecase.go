package usecase

import (
	"modern-go-app-structure/internal/domain"
	"modern-go-app-structure/internal/service"
)

// OrderUsecase - сценарии работы с заказами
type OrderUsecase interface {
	CreateOrder(userID string, amount int) (*domain.Order, error)
}

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	paymentServ service.PaymentService
}

func NewOrderUsecase(
	orderRepo domain.OrderRepository,
	paymentServ service.PaymentService,
) OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		paymentServ: paymentServ,
	}
}

func (uc *orderUsecase) CreateOrder(userID string, amount int) (*domain.Order, error) {
	// Создаем заказ
	order, err := uc.orderRepo.CreateOrder(userID, amount)
	if err != nil {
		return nil, err
	}
	
	// Обрабатываем платеж
	if err := uc.paymentServ.ProcessPayment(amount); err != nil {
		return nil, err
	}
	
	// Обновляем статус заказа
	// В реальном приложении здесь была бы дополнительная логика
	
	return order, nil
}
