package main

import (
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/postgres"
	"ward-stock-backend/internal/usecase"

	userhttp "ward-stock-backend/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := postgres.Connect()
	db.AutoMigrate(&domain.User{})

	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	r := gin.Default()

	userhttp.NewUserHandler(r, userUsecase)

	r.Run(":8080")
}
