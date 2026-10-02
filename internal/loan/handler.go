package loan

import (
	"errors"

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
	var req CreateLoanRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid request body",
			),
		)
	}

	loan, err := h.service.Create(c.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, ErrCustomerRequired),
			errors.Is(err, ErrPrincipalInvalid),
			errors.Is(err, ErrInterestInvalid),
			errors.Is(err, ErrTenorInvalid):

			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(
					fiber.StatusBadRequest,
					err.Error(),
				),
			)

		default:
			return c.Status(fiber.StatusInternalServerError).JSON(
				response.Error(
					fiber.StatusInternalServerError,
					err.Error(),
				),
			)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.Success(
			fiber.StatusCreated,
			"loan created successfully",
			loan,
		),
	)
}
