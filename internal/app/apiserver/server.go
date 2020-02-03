package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/nostereal/login-system/internal/app/models"
	"github.com/nostereal/login-system/internal/app/store"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	logrus.Debug("Server created.")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/ping", s.handlePing())
	s.router.HandleFunc("/login", s.handleLogIn()).Methods(http.MethodPost)
	s.router.HandleFunc("/signup", s.handleSignUp()).Methods(http.MethodPost)
}

func (s *server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Debug("ping request handled")
		fmt.Fprint(w, "Pong!")
	}
}

func (s *server) handleLogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creds := &model.User{}
		err := json.NewDecoder(r.Body).Decode(creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := s.store.User().FindByEmail(creds.Email)
		if err != nil {
			if err == store.ErrUserNotFound {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(creds.Password))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. read user credentials from request body
		// 2. check if user data valid
		// 3. register user using UserRepository.Create(user)
		// 4. if credentials are not valid => return http.StatusBadRequest
		//    if get an error when encrypt pass => return http.StatusInternalServerError

		creds := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(creds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.User().Create(creds); err != nil {
			if err == model.ErrUserCredentialsAreNotValid {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
