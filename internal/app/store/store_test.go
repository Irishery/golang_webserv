package store_test

import (
	"fmt"
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
		databaseURL = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
			config.POSTGRES_USER, config.POSTGRES_PASSWORD,
			config.DATABASE_HOST, config.DATABASE_PORT, config.POSTGRES_DB, config.SSL_MODE)
	}

	os.Exit(m.Run())
}
