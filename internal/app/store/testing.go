package store

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := NewConfig()
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v", err)
	}

	// config.DatabaseURL = databaseURL
	store := New(config)
	if err := store.Open(); err != nil {
		t.Fatal(err)
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := store.db.Exec((fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))); err != nil {
				t.Fatal(err)
			}
		}

		store.Close()
	}
}
