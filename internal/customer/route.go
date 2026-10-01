package customer

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {
	customers := router.Group("/customer")

	customers.Post("/", handler.Create)
}
