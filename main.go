package main

import (
	"ward-stock-backend/database"
	"ward-stock-backend/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/", handler.GetUsers)
		userRoutes.GET("/:id", handler.GetUser)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}

	r.Run(":8080")
}
