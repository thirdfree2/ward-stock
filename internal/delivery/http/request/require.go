package request

import (
	"net/http"

	"ward-stock-backend/internal/delivery/http/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RequireJSON(c *gin.Context, out any) bool {
	if err := c.ShouldBindJSON(out); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return false
	}

	if err := validate.Struct(out); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation failed", err)
		return false
	}

	return true
}
