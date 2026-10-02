package loan

import (
	"errors"

	"github.com/google/uuid"

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

func (h *Handler) Get(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid loan ID",
			),
		)
	}

	loan, err := h.service.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, ErrLoanNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				response.Error(
					fiber.StatusNotFound,
					err.Error(),
				),
			)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(
			response.Error(
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.Success(
			fiber.StatusOK,
			"loan retrieved successfully",
			loan,
		),
	)
}

func (h *Handler) List(c fiber.Ctx) error {
	var req ListLoanRequest

	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid query parameters",
			),
		)
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	loans, total, err := h.service.List(
		c.Context(),
		req,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.Error(
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		response.SuccessWithPagination(
			fiber.StatusOK,
			"loans retrieved successfully",
			loans,
			response.Pagination{
				Limit:  req.Limit,
				Offset: req.Offset,
				Total:  total,
			},
		),
	)
}
