package sqlstore_test

import (
	"os"
	"testing"
)

var (
	dbURL string
)

func TestMain(m *testing.M) {
	dbURL = os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost port=5432 user=postgres password=4321 dbname=loginsys_db sslmode=disable"
	}

	os.Exit(m.Run())
}
