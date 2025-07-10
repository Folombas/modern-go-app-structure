package domain

import (
	"fmt"
	"strconv"
)

// User - бизнес-сущность пользователя
type User struct {
	ID   string
	Name string
}

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(name string) *User
	FindByID(id string) (*User, error)
}

type userRepo struct {
	users map[string]*User
}

func NewUserRepository() UserRepository {
	return &userRepo{
		users: make(map[string]*User),
	}
}

func (r *userRepo) CreateUser(name string) *User {
	id := "user-" + strconv.Itoa(len(r.users)+1)
	user := &User{ID: id, Name: name}
	r.users[id] = user
	return user
}

func (r *userRepo) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
