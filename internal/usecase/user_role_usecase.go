package usecase

import "ward-stock-backend/internal/domain"

type userRoleUsecase struct {
	userRoleUC domain.UserRoleRepository
}

func NewUserRoleUsecase(userRoleUC domain.UserRoleRepository) domain.UserRoleRepository {
	return &userRoleUsecase{userRoleUC: userRoleUC}
}

func (u *userRoleUsecase) GetRolesByUserID(userID uint) (*domain.Role, error) {
	return u.userRoleUC.GetRolesByUserID(userID)
}
