package customer

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hndryno/go-lms-postgresql/internal/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c fiber.Ctx) error {
	var req CreateRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid request body",
			),
		)
	}

	customer, err := h.service.Create(
		c.Context(),
		req,
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.Success(
			fiber.StatusCreated,
			"customer created successfully",
			customer,
		),
	)
}
