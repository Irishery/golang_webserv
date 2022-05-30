package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var databaseURL string

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := NewTestConfig()
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v", err)
	}

	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=postgres password=pas277 host=localhost dbname=currency_info_test sslmode=disable"
	}

	os.Exit(m.Run())
}
