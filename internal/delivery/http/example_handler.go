package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/yourusername/go-skeleton/internal/model"
	"github.com/yourusername/go-skeleton/internal/usecase"
	"github.com/yourusername/go-skeleton/pkg/validator"
)

type ExampleHandler struct {
	useCase usecase.ExampleUseCase
}

func NewExampleHandler(useCase usecase.ExampleUseCase) *ExampleHandler {
	return &ExampleHandler{
		useCase: useCase,
	}
}

func (h *ExampleHandler) Create(c *fiber.Ctx) error {
	var req model.ExampleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid request body",
			nil,
		))
	}

	if err := validator.ValidateStruct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			err.Error(),
			nil,
		))
	}

	result, err := h.useCase.Create(c.Context(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.NewResponse(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		))
	}

	return c.Status(http.StatusCreated).JSON(model.NewResponse(
		http.StatusCreated,
		"success",
		"Example created successfully",
		result,
	))
}

func (h *ExampleHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid ID",
			nil,
		))
	}

	result, err := h.useCase.GetByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(model.NewResponse(
			http.StatusNotFound,
			"error",
			"Example not found",
			nil,
		))
	}

	return c.Status(http.StatusOK).JSON(model.NewResponse(
		http.StatusOK,
		"success",
		"Example retrieved successfully",
		result,
	))
}

func (h *ExampleHandler) GetAll(c *fiber.Ctx) error {
	var filter model.ExampleFilterRequest

	if err := c.QueryParser(&filter); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid query parameters",
			nil,
		))
	}

	results, meta, err := h.useCase.GetAll(c.Context(), &filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.NewResponse(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		))
	}

	return c.Status(http.StatusOK).JSON(model.NewPaginationResponse(
		http.StatusOK,
		"success",
		"Examples retrieved successfully",
		results,
		*meta,
	))
}

func (h *ExampleHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid ID",
			nil,
		))
	}

	var req model.ExampleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid request body",
			nil,
		))
	}

	if err := validator.ValidateStruct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			err.Error(),
			nil,
		))
	}

	result, err := h.useCase.Update(c.Context(), uint(id), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.NewResponse(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		))
	}

	return c.Status(http.StatusOK).JSON(model.NewResponse(
		http.StatusOK,
		"success",
		"Example updated successfully",
		result,
	))
}

func (h *ExampleHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.NewResponse(
			http.StatusBadRequest,
			"error",
			"Invalid ID",
			nil,
		))
	}

	if err := h.useCase.Delete(c.Context(), uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.NewResponse(
			http.StatusInternalServerError,
			"error",
			err.Error(),
			nil,
		))
	}

	return c.Status(http.StatusOK).JSON(model.NewResponse(
		http.StatusOK,
		"success",
		"Example deleted successfully",
		nil,
	))
}
