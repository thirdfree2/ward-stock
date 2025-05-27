package repository

import "ward-stock-backend/internal/domain"

type UserRepository interface {
	GetByID(id uint) (*domain.User, error)
	List() ([]domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error
}
