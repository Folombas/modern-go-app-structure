package rest

import (
	"context"
	"time"
	
	"github.com/Folombas/modern-go-app-structure/internal/repository"
	"github.com/Folombas/modern-go-app-structure/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // Импорт шаблонизатора
)

func SetupRouter(
	userRepo repository.UserRepository,
	orderUsecase usecase.OrderUsecase,
) *fiber.App {
	// Создаем движок шаблонов
	engine := html.New("./web/templates", ".html")
	
	app := fiber.New(fiber.Config{
		Views: engine, // Подключаем шаблонизатор
	})

	// Инициализация обработчиков
	webHandler := NewWebHandler(userRepo, orderUsecase)
	orderHandler := NewOrderHandler(orderUsecase)

	// Веб-роуты
	app.Get("/", webHandler.IndexPage)
	app.Get("/orders", webHandler.OrdersPage)
	app.Get("/users", webHandler.UsersPage)

	// API роуты
	api := app.Group("/api/v1")
	{
		api.Post("/orders", orderHandler.CreateOrder)
		api.Get("/orders/:id", orderHandler.GetOrder)
	}

	// Статические файлы
	app.Static("/static", "./web/static")

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}

// WebHandler обработчики веб-интерфейса
type WebHandler struct {
	userRepo     repository.UserRepository
	orderUsecase usecase.OrderUsecase
}

func NewWebHandler(
	userRepo repository.UserRepository,
	orderUsecase usecase.OrderUsecase,
) *WebHandler {
	return &WebHandler{
		userRepo:     userRepo,
		orderUsecase: orderUsecase,
	}
}

func (h *WebHandler) IndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Главная",
	}, "base") // Используем базовый шаблон
}

func (h *WebHandler) OrdersPage(c *fiber.Ctx) error {
	// Получаем заказы через usecase
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	orders, err := h.orderUsecase.GetRecentOrders(ctx, 10)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	
	return c.Render("orders", fiber.Map{
		"Title":  "Заказы",
		"Orders": orders,
	}, "base") // Используем базовый шаблон
}

func (h *WebHandler) UsersPage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Используем репозиторий напрямую
	users, err := h.userRepo.GetAllUsers(ctx)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	
	return c.Render("users", fiber.Map{
		"Title": "Пользователи",
		"Users": users,
	}, "base") // Используем базовый шаблон
}