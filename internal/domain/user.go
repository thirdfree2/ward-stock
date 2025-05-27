package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
	List() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
