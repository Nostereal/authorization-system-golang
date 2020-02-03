package sqlstore_test

import (
	"testing"

	model "github.com/nostereal/login-system/internal/app/models"
	"github.com/nostereal/login-system/internal/app/store/sqlstore"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dbURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser()
	err := s.User().Create(u)

	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dbURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@test.com"
	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	tu := model.TestUser()
	s.User().Create(tu)

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
