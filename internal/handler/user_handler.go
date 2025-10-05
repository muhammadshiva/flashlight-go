package handler

import (
	"net/http"
	"strconv"

	"flashlight-go/internal/dto"
	"flashlight-go/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid request", err))
		return
	}

	user, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to create user", err))
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse("User created successfully", user))
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse("User not found", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("User retrieved successfully", user))
}

func (h *UserHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	users, meta, err := h.userService.GetAll(c.Request.Context(), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to retrieve users", err))
		return
	}

	c.JSON(http.StatusOK, dto.PaginatedSuccessResponse("Users retrieved successfully", users, *meta))
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid request", err))
		return
	}

	user, err := h.userService.Update(c.Request.Context(), uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to update user", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("User updated successfully", user))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid ID", err))
		return
	}

	if err := h.userService.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse("Failed to delete user", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("User deleted successfully", nil))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse("Invalid request", err))
		return
	}

	response, err := h.userService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse("Login failed", err))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse("Login successful", response))
}
