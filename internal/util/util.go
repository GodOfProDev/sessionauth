package util

import (
	"github.com/godofprodev/sessionauth/internal/auth"
	"github.com/godofprodev/sessionauth/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetAuthenticatedUser(c *fiber.Ctx) *models.UserSession {
	value := c.Locals(auth.LocalsUserKey)
	if user, ok := value.(*models.UserSession); ok {
		return user
	}

	return nil
}
