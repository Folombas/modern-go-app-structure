package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Folombas/modern-go-app-structure/internal/domain"
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

// Close закрывает соединение с базой данных
func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

// CreateUser создает нового пользователя
func (r *PostgresRepository) CreateUser(ctx context.Context, name string) (*domain.User, error) {
	const query = `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id int // Получаем ID как число из БД
	err := r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &domain.User{
		ID:   fmt.Sprintf("%d", id), // Конвертируем в строку
		Name: name,
	}, nil
}

// GetUser возвращает пользователя по ID
func (r *PostgresRepository) GetUser(ctx context.Context, id string) (*domain.User, error) {
	const query = `SELECT name FROM users WHERE id = $1`
	var name string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &domain.User{ID: id, Name: name}, nil
}

// CreateOrder создает новый заказ
func (r *PostgresRepository) CreateOrder(ctx context.Context, userID string, amount int) (*domain.Order, error) {
	// Проверяем существование пользователя
	if _, err := r.GetUser(ctx, userID); err != nil {
		return nil, fmt.Errorf("invalid user: %w", err)
	}

	const query = `
		INSERT INTO orders (user_id, amount, status) 
		VALUES ($1, $2, 'created') 
		RETURNING id, created_at`
		
	var id int // Получаем ID как число из БД
	var createdAt time.Time
	
	err := r.db.QueryRowContext(ctx, query, userID, amount).Scan(&id, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	
	return &domain.Order{
		ID:        fmt.Sprintf("%d", id), // Конвертируем в строку
		UserID:    userID,
		Amount:    amount,
		Status:    "created",
		CreatedAt: createdAt,
	}, nil
}

// GetOrder возвращает заказ по ID
func (r *PostgresRepository) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	const query = `
		SELECT user_id, amount, status, created_at 
		FROM orders 
		WHERE id = $1`
		
	var userID string
	var amount int
	var status string
	var createdAt time.Time
	
	err := r.db.QueryRowContext(ctx, query, id).Scan(&userID, &amount, &status, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	
	return &domain.Order{
		ID:        id,
		UserID:    userID,
		Amount:    amount,
		Status:    status,
		CreatedAt: createdAt,
	}, nil
}

// Получаем последние заказы
func (r *PostgresRepository) GetRecentOrders(ctx context.Context, limit int) ([]*domain.Order, error) {
	const query = `
		SELECT id, user_id, amount, status, created_at 
		FROM orders 
		ORDER BY created_at DESC 
		LIMIT $1`
		
	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	defer rows.Close()
	
	var orders []*domain.Order
	for rows.Next() {
		var id int
		var userID string
		var amount int
		var status string
		var createdAt time.Time
		
		if err := rows.Scan(&id, &userID, &amount, &status, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		
		orders = append(orders, &domain.Order{
			ID:        fmt.Sprintf("%d", id),
			UserID:    userID,
			Amount:    amount,
			Status:    status,
			CreatedAt: createdAt,
		})
	}
	
	return orders, nil
}

// GetAllUsers возвращает всех пользователей
func (r *PostgresRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	const query = `SELECT id, name FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()
	
	var users []*domain.User
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &domain.User{
			ID:   fmt.Sprintf("%d", id),
			Name: name,
		})
	}
	return users, nil
}