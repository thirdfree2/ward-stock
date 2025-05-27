package domain

import (
	"ward-stock-backend/internal/domain/common"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Code  string `gorm:"uniqueIndex" json:"code"`
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
	common.BaseModel
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
	List() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
