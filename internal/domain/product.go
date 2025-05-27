package domain

import (
	"ward-stock-backend/internal/domain/common"
)

type Product struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	common.BaseModel
}

type ProductRepository interface {
	GetByID(id uint) (*Product, error)
	List() ([]Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
}
