package repayment

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	router fiber.Router,
	handler *Handler,
) {
	repayments := router.Group("/repayments")

	repayments.Post("/", handler.Create)
}
