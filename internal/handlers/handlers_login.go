package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"sessionauth/internal/models"
	"sessionauth/internal/response"
	"time"
)

func (h *Handlers) HandleLogin(c *fiber.Ctx) error {
	params := new(models.LoginUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if params.Username == "" {
		return response.ErrRequired("user")
	}
	if params.Password == "" {
		return response.ErrRequired("password")
	}

	user, err := h.store.GetUser(params.Username)
	if err != nil {
		return response.ErrNotFound(params.Username)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return response.ErrIncorrectPassword()
	}

	session, err := h.session.GenerateSession(user.ID.String())
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:     "session",
		Value:    session,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return response.SuccessMessage("successfully logged in as " + user.Username)
}
