package model

import (
	"errors"
	"regexp"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
}

var (
	// ErrUserCredentialsAreNotValid ...
	ErrUserCredentialsAreNotValid = errors.New("Users credentials are not valid")
)

// IsValid ...
func (u *User) IsValid() bool {
	rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	isPasswordCorrect := u.EncryptedPassword != "" || (len(u.Password) >= 8 && len(u.Password) <= 32)
	return len(u.Email) < 255 && rxEmail.MatchString(u.Email) && isPasswordCorrect
}

// EncryptPassword trying to encrypt user's password and returns error if it fails
func (u *User) EncryptPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Error while hashing password")
		return err
	}
	u.EncryptedPassword = string(hashedPass)

	return nil
}

// CompareHashAndPassword ...
func (u *User) CompareHashAndPassword(password string) {

}

// TestUser ...
func TestUser() *User {
	return &User{
		Email:    "user@test.com",
		Password: "test_pass",
	}
}

// TestUserWithInvalidPassword ...
func TestUserWithInvalidPassword() *User {
	return &User{
		Email:    "user2@test.com",
		Password: "short",
	}
}

// TestUserWithInvalidEmail ...
func TestUserWithInvalidEmail() *User {
	return &User{
		Email:    "usertest.com",
		Password: "test_password",
	}
}
