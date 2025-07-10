package domain

import (
	"fmt"
	"strconv"
	"time"
)

// Order - бизнес-сущность заказа
type Order struct {
	ID        string
	UserID    string
	Amount    int
	Status    string
	CreatedAt time.Time
}

// OrderRepository - интерфейс для работы с заказами
type OrderRepository interface {
	CreateOrder(userID string, amount int) (*Order, error)
	FindByID(id string) (*Order, error)
}

type orderRepo struct {
	orders map[string]*Order
}

func NewOrderRepository() OrderRepository {
	return &orderRepo{
		orders: make(map[string]*Order),
	}
}

func (r *orderRepo) CreateOrder(userID string, amount int) (*Order, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("invalid order amount")
	}
	
	id := "order-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	order := &Order{
		ID:        id,
		UserID:    userID,
		Amount:    amount,
		Status:    "created",
		CreatedAt: time.Now(),
	}
	r.orders[id] = order
	return order, nil
}

func (r *orderRepo) FindByID(id string) (*Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order not found")
	}
	return order, nil
}