package postgres

import (
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/postgres/model"

	"gorm.io/gorm"
)

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) domain.UserRoleRepository {
	return &userRoleRepository{db: db}
}

func (r *userRoleRepository) GetRolesByUserID(userID uint) (*domain.Role, error) {
	var userRole model.UserRoleModel
	if err := r.db.Preload("Role").
		Where("user_id = ?", userID).
		First(&userRole).Error; err != nil {
		return nil, err
	}

	return &userRole.Role, nil
}
