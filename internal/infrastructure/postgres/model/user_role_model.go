package model

import (
	"ward-stock-backend/internal/domain"
)

type UserRoleModel struct {
	ID     uint
	UserID uint
	RoleID uint
	Role   domain.Role `gorm:"foreignKey:RoleID"`
	User   domain.User `gorm:"foreignKey:UserID"`
}

func (UserRoleModel) TableName() string {
	return "user_roles"
}
