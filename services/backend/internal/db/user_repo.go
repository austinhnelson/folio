package db

import (
	"database/sql"
	"folio/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, username, password_hash, created_at)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		user.Email,
		user.Username,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, username, password_hash, created_at
		FROM users
		WHERE email = ?
	`

	u := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&u.Id,
		&u.Email,
		&u.Username,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return u, nil
}
