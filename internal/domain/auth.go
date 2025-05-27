package domain

type AuthUsecase interface {
	Login(email string) (string, error)
}
