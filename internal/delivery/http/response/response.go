package response

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  string      `json:"status"` // เช่น "success", "error"
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"` // payload
}

// Success ใช้ตอบสำเร็จ
func Success(c *gin.Context, data any) {
	c.JSON(200, ApiResponse{
		Status:  "success",
		Message: "success",
		Data:    data,
	})
}

// Error ใช้ตอบ error
func Error(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, ApiResponse{
		Status:  "error",
		Data:    gin.H{"error": err.Error()},
		Message: message,
	})
}
