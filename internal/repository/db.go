package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Folombas/modern-go-app-structure/internal/domain"
	"time"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// PostgresRepository реализация работы с PostgreSQL
type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return &PostgresRepository{db: db}, nil
}

// Реализация методов UserRepository
func (r *PostgresRepository) CreateUser(ctx context.Context, name string) (*domain.User, error) {
	const query = `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id string
	err := r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &domain.User{ID: id, Name: name}, nil
}

// Реализация методов OrderRepository
func (r *PostgresRepository) CreateOrder(ctx context.Context, userID string, amount int) (*domain.Order, error) {
	const query = `
		INSERT INTO orders (user_id, amount, status) 
		VALUES ($1, $2, 'created') 
		RETURNING id, created_at`
		
	var id string
	var createdAt time.Time
	
	err := r.db.QueryRowContext(ctx, query, userID, amount).Scan(&id, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	
	return &domain.Order{
		ID:        id,
		UserID:    userID,
		Amount:    amount,
		Status:    "created",
		CreatedAt: createdAt,
	}, nil
}