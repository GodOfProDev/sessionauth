package response

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Msg
}

func ErrCreating(a string) APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    fmt.Sprintf("there was an issue creating the %v", a),
	}
}

func ErrNotFound(a string) APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    fmt.Sprintf("%v not found", a),
	}
}

func ErrUpdating(a string) APIError {
	return APIError{
		Status: fiber.StatusNotFound,
		Msg:    fmt.Sprintf("there was an issue updating the %v", a),
	}
}

func ErrParsingParams() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "there was an issue parsing params",
	}
}

func ErrEncryptingPassword() APIError {
	return APIError{
		Status: fiber.StatusInternalServerError,
		Msg:    "there was an issue encrypting the password",
	}
}

func ErrIncorrectPassword() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "incorrect password",
	}
}

func ErrLoggingOut() APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    "there was an issue logging you out",
	}
}

func ErrUnauthorized() APIError {
	return APIError{
		Status: fiber.StatusUnauthorized,
		Msg:    "you are not authorized",
	}
}

func ErrValidating(message string) APIError {
	return APIError{
		Status: fiber.StatusBadRequest,
		Msg:    message,
	}
}
