package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		fmt.Println(token)
		// 🔐 TODO: decode JWT here → e.g., userID, role, permissions

		// mock ตัวอย่าง (ให้ทำจริงภายหลัง)
		c.Set("user_id", 1)
		c.Set("role", "admin")
		c.Set("permissions", []string{"user:create", "user:read", "user:delete"})

		c.Next()
	}
}
