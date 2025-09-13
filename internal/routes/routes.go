package routes

import (
	"go-web-crud/internal/config"
	"go-web-crud/internal/handler"
	"go-web-crud/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	cfg := config.LoadConfig()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
	app.Post("/register", handler.Register)
	app.Post("/login", func(c *fiber.Ctx) error {
		return handler.Login(c, cfg.JWTSecret)
	})

	api := app.Group("/api", middleware.AuthRequired(cfg.JWTSecret))
	api.Get("/user", handler.GetUserProfile)

	view := app.Group("/view")
	view.Get("/register", handler.RegisterView)
	view.Get("/login", handler.LoginView)
	view.Get("/profile", handler.ProfileView)
	view.Get("/", handler.IndexView)

}
