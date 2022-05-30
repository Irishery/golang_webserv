package store

import (
	"log"

	"github.com/Irishery/golang_webserv.git/internal/app/model"
)

// CurrencyRepository ...
type CurrencyRepository struct {
	store *Store
}

// Create ...
func (r *CurrencyRepository) Create(cur *model.Currency) (*model.Currency, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO currency (symbol, price, volume, last_trade) VALUES ($1, $2, $3, $4) RETURNING id",
		cur.Symbol,
		cur.Price,
		cur.Volume,
		cur.LastTrade,
	).Scan(&cur.ID); err != nil {
		return nil, err
	}

	return cur, nil
}

// CreateMany ...
func (r *CurrencyRepository) CreateMany(curArray []*model.Currency) ([]*model.Currency, error) {
	for _, cur := range curArray {
		_, err := r.FindBySymbol(cur.Symbol)
		if err != nil {
			_, err = r.Create(cur)
			if err != nil {
				log.Print(err)
			}
		} else {
			_, err = r.Update(*cur)
			if err != nil {
				log.Print(err)
			}
		}
	}

	log.Print("Data has been updated")

	return curArray, nil
}

// Update ...
func (r *CurrencyRepository) Update(sur model.Currency) (*model.Currency, error) {
	if err := r.store.db.QueryRow(
		"UPDATE currency SET price=$2, volume=$3, last_trade=$4 WHERE symbol=$1 RETURNING  price, volume, last_trade",
		sur.Symbol,
		sur.Price,
		sur.Volume,
		sur.LastTrade,
	).Scan(
		&sur.Price,
		&sur.Volume,
		&sur.LastTrade,
	); err != nil {
		log.Print(err)

		return nil, err
	}

	return &sur, nil
}

// FindBySymbol ...
func (r *CurrencyRepository) FindBySymbol(symbol string) (*model.Currency, error) {
	c := &model.Currency{}
	if err := r.store.db.QueryRow(
		"SELECT id, symbol, price, volume, last_trade FROM currency WHERE symbol = $1",
		symbol,
	).Scan(
		&c.ID,
		&c.Symbol,
		&c.Price,
		&c.Volume,
		&c.LastTrade,
	); err != nil {
		return nil, err
	}

	return c, nil
}

// GetAll ...
func (r *CurrencyRepository) GetAll() (map[string]model.CurrencyOutput, error) {
	rows, err := r.store.db.Query("SELECT * FROM currency")
	if err != nil {
		log.Print("Error ", err)
	}
	defer rows.Close()

	curArray := make(map[string]model.CurrencyOutput)

	for rows.Next() {
		var rawCur model.Currency
		if err := rows.Scan(&rawCur.ID, &rawCur.Symbol, &rawCur.Price, &rawCur.Volume,
			&rawCur.LastTrade); err != nil {
			return nil, err
		}

		curArray[rawCur.Symbol] = model.CurrencyOutput{
			Price:     rawCur.Price,
			Volume:    rawCur.Volume,
			LastTrade: rawCur.LastTrade,
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return curArray, nil
}
