package util

import (
	"github.com/gofiber/fiber/v2"
	"sessionauth/internal/auth"
	"sessionauth/internal/models"
)

func GetAuthenticatedUser(c *fiber.Ctx) *models.UserSession {
	value := c.Locals(auth.LocalsUserKey)
	if user, ok := value.(*models.UserSession); ok {
		return user
	}

	return nil
}
