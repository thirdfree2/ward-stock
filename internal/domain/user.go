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
	GetByEmail(email string) (*User, error)
	GetPermissions(userID uint) ([]string, error)
	List() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type UserUsecase interface {
	GetUserByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	ListUsers() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}
