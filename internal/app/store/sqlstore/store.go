package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
	"github.com/nostereal/login-system/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{store: s}
	}

	return s.userRepository
}
