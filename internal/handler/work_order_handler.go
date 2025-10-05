package handler

import (
	"net/http"
	"strconv"

	"flashlight-go/internal/dto"
	"flashlight-go/internal/service"

	"github.com/gin-gonic/gin"
)

type WorkOrderHandler struct {
	workOrderService *service.WorkOrderService
}

func NewWorkOrderHandler(workOrderService *service.WorkOrderService) *WorkOrderHandler {
	return &WorkOrderHandler{workOrderService: workOrderService}
}

func (h *WorkOrderHandler) Create(c *gin.Context) {
	var req dto.CreateWorkOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid request", err))
		return
	}

	// Get cashier user ID from context (set by auth middleware)
	var cashierUserID *uint
	if userID, exists := c.Get("user_id"); exists {
		uid := userID.(uint)
		cashierUserID = &uid
	}

	workOrder, err := h.workOrderService.Create(c.Request.Context(), req, cashierUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to create work order", err))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse("Work order created successfully", workOrder))
}

func (h *WorkOrderHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	workOrder, err := h.workOrderService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse("Work order not found", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Work order retrieved successfully", workOrder))
}

func (h *WorkOrderHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	// Check if status filter is provided
	if status := c.Query("status"); status != "" {
		workOrders, err := h.workOrderService.GetByStatus(c.Request.Context(), status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to retrieve work orders", err))
			return
		}
		c.JSON(http.StatusOK, dto.SuccessResponse("Work orders retrieved successfully", workOrders))
		return
	}

	workOrders, meta, err := h.workOrderService.GetAll(c.Request.Context(), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to retrieve work orders", err))
		return
	}

	c.JSON(http.StatusOK, dto.PaginatedSuccessResponse("Work orders retrieved successfully", workOrders, *meta))
}

func (h *WorkOrderHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	var req dto.UpdateWorkOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid request", err))
		return
	}

	workOrder, err := h.workOrderService.Update(c.Request.Context(), uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to update work order", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Work order updated successfully", workOrder))
}

func (h *WorkOrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	if err := h.workOrderService.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to delete work order", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Work order deleted successfully", nil))
}
