package http

import (
	"ward-stock-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes adds all HTTP routes
func RegisterRoutes(r *gin.Engine, userRoleUC domain.UserRoleUsecase, userUC domain.UserUsecase, authUC domain.AuthUsecase) {
	NewUserHandler(r, userUC, userRoleUC)
	NewAuthHandler(r, authUC)
	// เพิ่ม handler อื่นที่นี่ เช่น NewProductHandler(r, ...)
}
