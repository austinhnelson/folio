package auth

import "folio/internal/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}
