package store

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config             *Config
	db                 *sql.DB
	currencyRepository *CurrencyRepository
}

// New ..
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	var databaseURL string = fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		s.config.POSTGRES_USER, s.config.POSTGRES_PASSWORD,
		s.config.DATABASE_HOST, s.config.POSTGRES_DB, s.config.SSL_MODE)

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// Currency ...
func (s *Store) Currency() *CurrencyRepository {
	if s.currencyRepository != nil {
		return s.currencyRepository
	}

	s.currencyRepository = &CurrencyRepository{
		store: s,
	}

	return s.currencyRepository
}
