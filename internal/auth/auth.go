package auth

import (
	"github.com/gofiber/fiber/v2"
	"sessionauth/internal/models"
	"sessionauth/internal/response"
	"sessionauth/internal/session"
	"sessionauth/internal/storage"
)

type Auth struct {
	storage storage.Storage
	session session.Session
}

func NewAuth(storage storage.Storage, session session.Session) *Auth {
	return &Auth{
		storage: storage,
		session: session,
	}
}

const LocalsUserKey = "user"

func (a *Auth) Authenticate(c *fiber.Ctx) error {
	c.Locals(LocalsUserKey, nil)
	cookie := c.Cookies("session")

	if cookie == "" || len(cookie) < 8 || cookie[:7] != "Bearer " {
		return c.JSON(fiber.Map{"error": "invalid session header"})
	}

	// get the session id
	sessionId := cookie[7:]

	userID, err := a.session.GetUserIDBySession(sessionId)
	if err != nil {
		return err
	}

	if userID != "" {
		return response.ErrUnauthorized()
	}

	user, err := a.storage.GetUser(userID)
	if err != nil {
		return err
	}

	userSession := models.UserSession{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	c.Locals(LocalsUserKey, &userSession)

	return c.Next()
}
