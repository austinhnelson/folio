package auth

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthService AuthService
}

type AuthService interface {
	Register(email, username, password string) (string, error)
	Login(email, password string) (string, error)
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := authHandler.AuthService.Register(req.Email, req.Username, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // false since local dev. in prod, use true.
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   900, // 15 minutes
	})

	w.WriteHeader(http.StatusOK)
}

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := authHandler.AuthService.Login(req.Email, req.Password)

	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // false since local dev. in prod, use true.
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   900, // 15 minutes
	})

	w.WriteHeader(http.StatusOK)
}
