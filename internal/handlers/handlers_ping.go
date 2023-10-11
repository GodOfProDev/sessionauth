package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePing(c *fiber.Ctx) error {
	return c.SendString("ping")
}
