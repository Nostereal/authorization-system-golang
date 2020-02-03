package sqlstore

import (
	"database/sql"

	model "github.com/nostereal/login-system/internal/app/models"
	"github.com/nostereal/login-system/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create adds user (also encrypt password) to database (if it's valid)
func (r *UserRepository) Create(u *model.User) error {
	if !u.IsValid() {
		return model.ErrUserCredentialsAreNotValid
	}

	err := u.EncryptPassword()
	if err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail finds user in the database by email and returns User struct if there is no errors. Otherwise, return nil, error.
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}
