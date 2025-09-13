package handler

import "github.com/gofiber/fiber/v2"

func IndexView(c *fiber.Ctx) error {
	return c.SendFile("./internal/views/index.html")
}

func RegisterView(c *fiber.Ctx) error {
	return c.SendFile("./internal/views/register.html")
}

func LoginView(c *fiber.Ctx) error {
	return c.SendFile("./internal/views/login.html")
}

func ProfileView(c *fiber.Ctx) error {
	return c.SendFile("./internal/views/profile.html")
}
