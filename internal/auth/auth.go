package auth

import (
	"github.com/gofiber/fiber/v2"
	"sessionauth/internal/response"
	"sessionauth/internal/session"
)

type Auth struct {
	session session.Session
}

func NewAuth(session session.Session) *Auth {
	return &Auth{
		session: session,
	}
}

func (a *Auth) Authenticate(c *fiber.Ctx) error {
	sessionHeader := c.Get("Authorization")

	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := sessionHeader[7:]

	user, err := a.session.GetUserBySession(sessionId)
	if err != nil {
		return err
	}

	if user != "" {
		return response.ErrUnauthorized()
	}

	return c.Next()
}
