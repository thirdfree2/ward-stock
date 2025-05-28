package dto

import (
	"ward-stock-backend/internal/domain"
)

func ToDomainUser(req CreateUserRequest) domain.User {
	return domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func ToUserResponse(u *domain.User, r *domain.Role) *UserResponse {
	return &UserResponse{
		ID:    u.ID,
		Code:  u.Code,
		Name:  u.Name,
		Email: u.Email,
		Role:  *r,
	}
}
