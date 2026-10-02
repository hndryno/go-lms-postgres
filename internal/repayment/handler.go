package repayment

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
	var req CreateRepaymentRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid request body",
			),
		)
	}

	err := h.service.Create(c.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, ErrLoanRequired),
			errors.Is(err, ErrAmountInvalid):

			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(
					fiber.StatusBadRequest,
					err.Error(),
				),
			)

		case errors.Is(err, ErrLoanNotFound):
			return c.Status(fiber.StatusNotFound).JSON(
				response.Error(
					fiber.StatusNotFound,
					err.Error(),
				),
			)

		case errors.Is(err, ErrAmountExceedsRemaining):
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
			"repayment created successfully",
			nil,
		),
	)
}
