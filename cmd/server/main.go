package main

import (
	"ward-stock-backend/internal/delivery/http"
	"ward-stock-backend/internal/delivery/http/middleware"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/postgres"
	"ward-stock-backend/internal/infrastructure/service"
	"ward-stock-backend/internal/usecase"

	_ "ward-stock-backend/docs"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := postgres.Connect()
	db.AutoMigrate(
		&domain.User{},
		&domain.RunningNumber{},
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.UserRole{},
	)

	userRepo := postgres.NewUserRepository(db)
	runningRepo := service.NewRunningNumberService(db)

	userUsecase := usecase.NewUserUsecase(userRepo, runningRepo)

	r := gin.Default()

	http.RegisterRoutes(r, userUsecase)
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.AuthMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
