package auth

import (
	"backend/internal/models"
	"folio/internal/models"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

type AuthService struct {
	userRepo UserRepo
	jwtToken string
}

func NewAuthService(userRepo UserRepo, jwtToken string) *AuthService {
	return
}
