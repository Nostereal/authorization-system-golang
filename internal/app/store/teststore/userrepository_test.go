package teststore_test

import (
	"testing"

	model "github.com/nostereal/login-system/internal/app/models"
	"github.com/nostereal/login-system/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser_CorrectCredentials(t *testing.T) {
	s := teststore.New()
	u := model.TestUser()
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

	u = model.TestUserWithInvalidEmail()
	assert.Error(t, s.User().Create(u))

	u = model.TestUserWithInvalidPassword()
	assert.Error(t, s.User().Create(u))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	u := model.TestUser()
	s.User().Create(u)
	fndUser, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, fndUser)

	u = model.TestUserWithInvalidEmail()
	fndUser, err = s.User().FindByEmail(u.Email)
	assert.Error(t, err)
	assert.Nil(t, fndUser)
}
