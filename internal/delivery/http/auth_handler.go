package http

import (
	"net/http"
	"ward-stock-backend/internal/delivery/http/dto"
	"ward-stock-backend/internal/delivery/http/response"
	"ward-stock-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc domain.AuthUsecase) {
	h := &AuthHandler{usecase: uc}

	users := r.Group("/auth")
	{
		users.POST("/", h.Login)
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	token, err := h.usecase.Login(req.Email)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	response.Success(c, dto.LoginResponse{Token: token})
}
