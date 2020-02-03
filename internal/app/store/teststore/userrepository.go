package teststore

import (
	"log"

	model "github.com/nostereal/login-system/internal/app/models"
	"github.com/nostereal/login-system/internal/app/store"
)

// UserRepository is fake implementation of original UserRepository.
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create fake user
func (r *UserRepository) Create(u *model.User) error {
	if !u.IsValid() {
		return model.ErrUserCredentialsAreNotValid
	}

	if err := u.EncryptPassword(); err != nil {
		return err
	}

	log.Printf("Inside testCreate: %s\n", u.EncryptedPassword)

	u.ID = len(r.users)
	r.users[u.Email] = u

	return nil
}

// FindByEmail is a fake implementation that checks if user with certain email exists in users map
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrUserNotFound
	}

	return u, nil
}
