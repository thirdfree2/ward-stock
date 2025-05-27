package dto

import "ward-stock-backend/internal/domain"

func ToDomainUser(req CreateUserRequest) domain.User {
	return domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
}
