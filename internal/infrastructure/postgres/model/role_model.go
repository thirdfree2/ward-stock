package model

import "ward-stock-backend/internal/domain"

type RoleModel struct {
	ID   uint
	Name string
}

func (RoleModel) TableName() string {
	return "roles"
}

func ToDomainRole(r *RoleModel) *domain.Role {
	return &domain.Role{
		ID:   r.ID,
		Name: r.Name,
	}
}
