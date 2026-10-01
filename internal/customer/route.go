package customer

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {
	customers := router.Group("/customer")

	customers.Post("/", handler.Create)
	customers.Get("/", handler.List)
	customers.Get("/:id", handler.Get)
	customers.Put("/:id", handler.Update)
	// customers.Delete("/:id", handler.Delete)
}
