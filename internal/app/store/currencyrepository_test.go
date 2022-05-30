package store_test

import (
	"testing"

	"github.com/Irishery/golang_webserv.git/internal/app/model"
	"github.com/Irishery/golang_webserv.git/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currency")

	c, err := s.Currency().Create(&model.Currency{
		Symbol:    "test",
		Price:     0.1,
		Volume:    0.2,
		LastTrade: 0.3,
	})

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCurrencyRepository_FindBySymbol(t *testing.T) {
	store, teardown := store.TestStore(t, databaseURL)
	defer teardown("currency")

	symbol := "test"
	_, err := store.Currency().FindBySymbol(symbol)
	assert.Error(t, err)

	store.Currency().Create(&model.Currency{
		Symbol:    "test",
		Price:     0.1,
		Volume:    0.2,
		LastTrade: 0.3,
	})

	c, err := store.Currency().FindBySymbol(symbol)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
