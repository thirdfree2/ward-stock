package main

import (
	"log"
	"ward-stock-backend/internal/auth"
	"ward-stock-backend/internal/delivery/http"
	"ward-stock-backend/internal/delivery/http/middleware"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/postgres"
	"ward-stock-backend/internal/infrastructure/service"
	"ward-stock-backend/internal/usecase"

	_ "ward-stock-backend/cmd/server/docs"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := postgres.Connect()
	db.AutoMigrate(
		&domain.User{},
		&domain.RunningNumber{},
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.UserRole{},
	)
	jwtSvc := service.NewJwtService()
	userRepo := postgres.NewUserRepository(db)
	userRoleRepo := postgres.NewUserRoleRepository(db)
	runningRepo := service.NewRunningNumberService(db)
	userUsecase := usecase.NewUserUsecase(userRepo, runningRepo)
	userRoleUsecase := usecase.NewUserRoleUsecase(userRoleRepo)
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtSvc)
	r := gin.Default()

	http.RegisterRoutes(r, userRoleUsecase, userUsecase, authUsecase)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.LoggerMiddleware())
	r.Use(auth.AuthMiddleware())

	r.Run(":8080")
}
