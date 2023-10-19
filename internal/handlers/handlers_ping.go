package handlers

import (
	"github.com/godofprodev/sessionauth/internal/response"
	"github.com/godofprodev/sessionauth/internal/util"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePing(c *fiber.Ctx) error {
	user := util.GetAuthenticatedUser(c)
	if user == nil {
		return response.ErrUnauthorized()
	}

	return c.SendString("pinged by " + user.Username)
}
