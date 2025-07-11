package rest

import (
	"time"
	
	"github.com/Folombas/modern-go-app-structure/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type WebHandler struct {
	userRepo  repository.UserRepository
	orderRepo repository.OrderRepository
}

func NewWebHandler(
	userRepo repository.UserRepository,
	orderRepo repository.OrderRepository,
) *WebHandler {
	return &WebHandler{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

func (h *WebHandler) IndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Главная",
	})
}

func (h *WebHandler) OrdersPage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	orders, err := h.orderRepo.GetRecentOrders(ctx, 10)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	
	return c.Render("orders", fiber.Map{
		"Title":  "Заказы",
		"Orders": orders,
	})
}

func (h *WebHandler) UsersPage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	users, err := h.userRepo.GetAllUsers(ctx)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	
	return c.Render("users", fiber.Map{
		"Title":  "Пользователи",
		"Users":  users,
	})
}