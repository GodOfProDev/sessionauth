package auth

import (
	"github.com/godofprodev/sessionauth/internal/models"
	"github.com/godofprodev/sessionauth/internal/response"
	"github.com/godofprodev/sessionauth/internal/session"
	"github.com/godofprodev/sessionauth/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	if cookie == "" || !validateUUID(cookie) {
		return c.JSON(fiber.Map{"error": "invalid session"})
	}

	sessionId := cookie

	userID, err := a.session.GetUserIDBySession(sessionId)
	if err != nil {
		return err
	}

	if userID == "" {
		return response.ErrUnauthorized()
	}

	user, err := a.storage.GetUserByID(userID)
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

func validateUUID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	}

	return true
}
