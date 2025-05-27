package http

import (
	"ward-stock-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes adds all HTTP routes
func RegisterRoutes(r *gin.Engine, userUC usecase.UserUsecase, authUC usecase.AuthUsecase) {
	NewUserHandler(r, userUC)
	NewAuthHandler(r, authUC)
	// เพิ่ม handler อื่นที่นี่ เช่น NewProductHandler(r, ...)
}
