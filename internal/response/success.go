package response

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type APISuccessData struct {
	Status int `json:"-"`
	Data   any `json:"-"`
}

type APISuccessResponse struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APISuccessData) Error() string {
	jsonBytes, _ := json.Marshal(e.Data)
	return string(jsonBytes)
}

func (e APISuccessResponse) Error() string {
	return e.Msg
}

func SuccessCreated(a any) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusCreated,
		Data:   a,
	}
}

func SuccessGotten(a any) APISuccessData {
	return APISuccessData{
		Status: fiber.StatusOK,
		Data:   a,
	}
}

func SuccessMessage(a string) APISuccessResponse {
	return APISuccessResponse{
		Status: fiber.StatusOK,
		Msg:    a,
	}
}
