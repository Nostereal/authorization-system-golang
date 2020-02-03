package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/nostereal/login-system/internal/app/models"
	storepkg "github.com/nostereal/login-system/internal/app/store"
	"github.com/nostereal/login-system/internal/app/store/sqlstore"
	"github.com/nostereal/login-system/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAPIServer_HandlePing(t *testing.T) {
	db, err := newDB("host=localhost port=5432 user=postgres password=4321 dbname=loginsys_db sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	srv.handlePing().ServeHTTP(rec, req)
	assert.Equal(t, "Pong!", rec.Body.String())
}

func TestAPIServer_HandleLogIn(t *testing.T) {
	store := teststore.New()
	// check correct status 200
	u := model.TestUser()
	err := store.User().Create(u)
	if err != nil {
		t.Fatal(err)
	}

	srv := newServer(store)
	rec := httptest.NewRecorder()

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(u); err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/login", buf)
	srv.handleLogIn().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// check correct error
	buf = new(bytes.Buffer)
	u = model.TestUser()
	u.Email = "new@email.com"
	if err := json.NewEncoder(buf).Encode(u); err != nil {
		t.Fatal(err)
	}

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/login", buf)
	srv.handleLogIn().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAPIServer_HandleSignUp(t *testing.T) {
	store := teststore.New()
	srv := newServer(store)

	validUsr := model.TestUser()
	invalidUsr := model.TestUserWithInvalidPassword()

	// Test case 1: pass user with correct credentials and check if it's successfully signed up
	rec := httptest.NewRecorder()
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(validUsr)

	req, _ := http.NewRequest(http.MethodPost, "/signup", buf)
	srv.handleSignUp().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	// check if this user registered correctly
	foundUsr, err := store.User().FindByEmail(validUsr.Email)
	assert.NoError(t, err)
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(foundUsr.EncryptedPassword), []byte(validUsr.Password)))

	// Test case 2: pass user with incorrect credentials and make sure about correct response status and that user isn't registered
	buf = new(bytes.Buffer)
	json.NewEncoder(buf).Encode(invalidUsr)
	req, _ = http.NewRequest(http.MethodPost, "/signup", buf)
	rec = httptest.NewRecorder()

	srv.handleSignUp().ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	fndUsr, err := store.User().FindByEmail(invalidUsr.Email)
	assert.EqualError(t, err, storepkg.ErrUserNotFound.Error())
	assert.Nil(t, fndUsr)
}
