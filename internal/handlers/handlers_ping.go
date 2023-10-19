package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sessionauth/internal/response"
	"sessionauth/internal/util"
)

func (h *Handlers) HandlePing(c *fiber.Ctx) error {
	user := util.GetAuthenticatedUser(c)
	if user != nil {
		return response.ErrUnauthorized()
	}

	return c.SendString("pinged by " + user.Username)
}
