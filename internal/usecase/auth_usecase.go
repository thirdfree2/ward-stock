package usecase

import (
	"errors"
	"ward-stock-backend/internal/domain"
	"ward-stock-backend/internal/infrastructure/service"
)

type authUsecase struct {
	userRepo domain.UserRepository
	jwtSvc   service.JwtService
}

func NewAuthUsecase(userRepo domain.UserRepository, jwtSvc service.JwtService) domain.AuthUsecase {
	return &authUsecase{userRepo, jwtSvc}
}

func (u *authUsecase) Login(email string) (string, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	perms, err := u.userRepo.GetPermissions(user.ID)
	if err != nil {
		return "", errors.New("Can't Get Permission")
	}
	// 👈 อาจ preload จาก role/permission
	return u.jwtSvc.GenerateToken(user.ID, user.Email, perms)
}
