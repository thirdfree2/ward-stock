package dto

import "ward-stock-backend/internal/domain"

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID    uint        `json:"id"`
	Code  string      `json:"code"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
}
