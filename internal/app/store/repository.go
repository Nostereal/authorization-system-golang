package store

import model "github.com/nostereal/login-system/internal/app/models"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
