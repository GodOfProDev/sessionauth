package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"sessionauth/internal/auth"
	"sessionauth/internal/handlers"
	"sessionauth/internal/response"
	"sessionauth/internal/session"
	"sessionauth/internal/storage"
)

type Router struct {
	app     *fiber.App
	store   storage.Storage
	session session.Session
}

func New(store storage.Storage, session session.Session) *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	return &Router{
		app:     app,
		store:   store,
		session: session,
	}
}

func (r *Router) Listen() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	err := r.app.Listen(fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RegisterMiddlewares() {
	r.app.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] ${white}| ${status} | ${latency} | ${white}${ip} | ${method} | ${white}${path}\n",
		TimeZone: "UTC",
	}))
	r.app.Use(cors.New(cors.Config{AllowCredentials: true}))
}

func (r *Router) RegisterHandlers() {
	h := handlers.New(r.store)
	authMiddleware := auth.NewAuth(r.session)

	v1 := r.app.Group("/v1")

	v1.Get("/ping", authMiddleware.Authenticate, h.HandlePing)
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case response.APIError:
		return c.Status(e.Status).JSON(e)
	case response.APISuccessData:
		return c.Status(e.Status).JSON(e.Data)
	case response.APISuccessResponse:
		return c.Status(e.Status).JSON(e)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"message": "internal server error"})
	}
}
