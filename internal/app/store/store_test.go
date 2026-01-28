package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://admin:admin@localhost:5432/postgres?sslmode=disable"
	}
	os.Exit(m.Run())

}
