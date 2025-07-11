package repository

import (
	"context"
	"github.com/Folombas/modern-go-app-structure/internal/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, userID string, amount int) (*domain.Order, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, name string) (*domain.User, error)
}