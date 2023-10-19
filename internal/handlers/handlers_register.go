package handlers

import (
	"github.com/godofprodev/sessionauth/internal/models"
	"github.com/godofprodev/sessionauth/internal/response"
	"github.com/godofprodev/sessionauth/internal/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) HandleRegister(c *fiber.Ctx) error {
	params := new(models.RegisterUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if err := h.validator.Validate(params); err != nil {
		return validator.FormatValidationErrors(err)
	}

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
	if err != nil {
		return response.ErrEncryptingPassword()
	}

	user := models.User{
		ID:       uuid.New(),
		Username: params.Username,
		Email:    params.Email,
		Password: string(encryptedPass),
	}

	if err := h.store.CreateUser(&user); err != nil {
		return response.ErrCreating("user")
	}

	return response.SuccessMessage("successfully created your account")
}
