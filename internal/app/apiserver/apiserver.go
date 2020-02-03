package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/nostereal/login-system/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	logrus.Info("Starting server...")

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logrus.Error("Error while opening db connection: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.Error("Error while pinging db: ", err)
		return nil, err
	}

	logrus.Info("Database started successfully!")

	return db, nil
}
