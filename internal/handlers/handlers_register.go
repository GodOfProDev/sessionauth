package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"sessionauth/internal/models"
	"sessionauth/internal/response"
)

func (h *Handlers) HandleRegister(c *fiber.Ctx) error {
	params := new(models.RegisterUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if params.Username == "" {
		return response.ErrRequired("user")
	}
	if params.Password == "" {
		return response.ErrRequired("password")
	}

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
	if err != nil {
		return response.ErrEncryptingPassword()
	}

	user := models.User{
		ID:       uuid.New(),
		Username: params.Username,
		Password: string(encryptedPass),
	}

	if err := h.store.CreateUser(&user); err != nil {
		return response.ErrCreating("user")
	}

	return response.SuccessCreated(user)
}
