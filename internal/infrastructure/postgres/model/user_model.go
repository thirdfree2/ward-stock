package model

import (
	"ward-stock-backend/internal/domain"
)

type UserModel struct {
	ID    uint
	Code  string
	Name  string
	Email string
	Roles []domain.UserRole `gorm:"foreignKey:UserID"`
}

func ToDomainUser(u *UserModel) *domain.User {
	return &domain.User{
		ID:    u.ID,
		Code:  u.Code,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (UserModel) TableName() string {
	return "users"
}
