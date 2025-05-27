package dto

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
