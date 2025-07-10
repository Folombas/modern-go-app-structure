package service

import "errors"

// PaymentService - сервис обработки платежей
type PaymentService interface {
	ProcessPayment(amount int) error
}

type paymentService struct{}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (s *paymentService) ProcessPayment(amount int) error {
	if amount <= 0 {
		return errors.New("invalid payment amount")
	}
	
	// Здесь была бы реальная логика обработки платежа
	// Например, интеграция с платежным шлюзом
	return nil
}