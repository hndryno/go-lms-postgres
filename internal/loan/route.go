package loan

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {
	loans := router.Group("/loans")

	loans.Post("/", handler.Create)
	loans.Get("/", handler.List)
	loans.Get("/:id", handler.Get)
}
