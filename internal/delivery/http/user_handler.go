package http

import (
	"net/http"
	"strconv"

	"ward-stock-backend/internal/delivery/http/dto"
	"ward-stock-backend/internal/delivery/http/request"
	"ward-stock-backend/internal/delivery/http/response"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, uc usecase.UserUsecase) {
	h := &UserHandler{usecase: uc}

	users := r.Group("/users")
	{
		users.GET("/", h.ListUsers)
		users.GET("/:id", h.GetUserByID)
		users.POST("/", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponse
// @Failure 404 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.usecase.GetUserByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.usecase.ListUsers()
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err)
		return
	}
	response.Success(c, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param body body domain.User true "User data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if !request.RequireJSON(c, &req) {
		return
	}
	user := dto.ToDomainUser(req)
	if err := h.usecase.CreateUser(&user); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	response.Success(c, user)
}

// UpdateUser godoc
// @Summary Update user data
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body domain.User true "User data to update"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}
	user.ID = uint(id)
	if err := h.usecase.UpdateUser(&user); err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err)
		return
	}
	response.Success(c, user)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteUser(uint(id)); err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	response.Success(c, "")
}
