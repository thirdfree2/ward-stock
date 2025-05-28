package usecase

import (
	"ward-stock-backend/internal/domain"
)

type userUsecase struct {
	repo        domain.UserRepository
	runningRepo domain.RunningNumberRepository
}

func NewUserUsecase(repo domain.UserRepository, runningRepo domain.RunningNumberRepository) domain.UserUsecase {
	return &userUsecase{repo: repo, runningRepo: runningRepo}
}

func (u *userUsecase) GetByEmail(email string) (*domain.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userUsecase) GetUserByID(id uint) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) ListUsers() ([]domain.User, error) {
	return u.repo.List()
}

func (u *userUsecase) CreateUser(user *domain.User) error {
	code, err := u.runningRepo.NextNumber("USER", "USR", 5)
	if err != nil {
		return err
	}

	user.Code = code

	return u.repo.Create(user)
}

func (u *userUsecase) UpdateUser(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}
