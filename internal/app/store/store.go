package store

import "errors"

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("User not found")
)

// Store ...
type Store interface {
	User() UserRepository
}
