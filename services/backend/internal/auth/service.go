package auth

import (
	"folio/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo  UserRepo
	jwtSecret string
}

func NewAuthService(userRepo UserRepo, secret string) *authService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: secret,
	}
}

func (a *authService) Register(email, username, password string) (string, error) {
	hash, passGenErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if passGenErr != nil {
		return "", passGenErr
	}

	userToCreate := models.User{
		ID:        uuid.NewString(),
		Email:     email,
		Username:  username,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	createUserErr := a.userRepo.CreateUser(&userToCreate)

	if createUserErr != nil {
		return "", createUserErr
	}

	return a.generateJwt(userToCreate.ID, userToCreate.Email)
}

func (a *authService) Login(email, password string) (string, error) {
	user, err := a.userRepo.GetByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	return a.generateJwt(user.ID, user.Email)
}

func (a *authService) generateJwt(userID string, email string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add((15 * time.Minute))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "folio-backend",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}
