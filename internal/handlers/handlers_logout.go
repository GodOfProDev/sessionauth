package handlers

import (
	"github.com/godofprodev/sessionauth/internal/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleLogout(c *fiber.Ctx) error {
	sessionId := c.Cookies("session")

	err := h.session.DeleteSession(sessionId)
	if err != nil {
		return response.ErrLoggingOut()
	}

	return response.SuccessMessage("successfully logged out")
}
