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
		databaseURL = "user=postgres password=pas277 host=localhost dbname=currency_info_test sslmode=disable"

	}

	os.Exit(m.Run())
}
