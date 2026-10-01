package customer

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

func (h *Handler) List(c fiber.Ctx) error {
	limit := int32(10)
	offset := int32(0)

	if value := c.Query("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(
					fiber.StatusBadRequest,
					"invalid limit",
				),
			)
		}

		limit = int32(parsed)
	}

	if value := c.Query("offset"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(
					fiber.StatusBadRequest,
					"invalid offset",
				),
			)
		}

		offset = int32(parsed)
	}

	customers, total, err := h.service.List(
		c.Context(),
		limit,
		offset,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.Error(
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return c.JSON(
		response.SuccessWithPagination(
			fiber.StatusOK,
			"customers retrieved successfully",
			customers,
			response.Pagination{
				Limit:  int(limit),
				Offset: int(offset),
				Total:  total,
			},
		),
	)
}

func (h *Handler) Get(c fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid customer id",
			),
		)
	}

	customer, err := h.service.Get(
		c.Context(),
		id,
	)
	if err != nil {
		if errors.Is(err, ErrCustomerNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				response.Error(
					fiber.StatusNotFound,
					"customer not found",
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

	return c.JSON(
		response.Success(
			fiber.StatusOK,
			"customer retrieved successfully",
			customer,
		),
	)
}

func (h *Handler) Update(c fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid customer id",
			),
		)
	}

	var req UpdateCustomerRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error(
				fiber.StatusBadRequest,
				"invalid request body",
			),
		)
	}

	customer, err := h.service.Update(
		c.Context(),
		id,
		req,
	)
	if err != nil {
		if errors.Is(err, ErrCustomerNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				response.Error(
					fiber.StatusNotFound,
					"customer not found",
				),
			)
		}

		if errors.Is(err, ErrNameRequired) ||
			errors.Is(err, ErrEmailRequired) {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.Error(
					fiber.StatusBadRequest,
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

	return c.JSON(
		response.Success(
			fiber.StatusOK,
			"customer updated successfully",
			customer,
		),
	)
}
