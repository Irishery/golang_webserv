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
func (r *CurrencyRepository) Create(c *model.Currency) (*model.Currency, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO currency (symbol, price, volume, last_trade) VALUES ($1, $2, $3, $4) RETURNING id",
		c.Symbol,
		c.Price,
		c.Volume,
		c.LastTrade,
	).Scan(&c.ID); err != nil {
		return nil, err
	}

	return c, nil
}

// CreateMany ...
func (r *CurrencyRepository) CreateMany(cur_array []*model.Currency) ([]*model.Currency, error) {
	for _, cur := range cur_array {
		_, err := r.FindBySymbol(cur.Symbol)
		if err != nil {
			log.Print(err)
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

	return cur_array, nil
}

// Update ...
func (r *CurrencyRepository) Update(c model.Currency) (*model.Currency, error) {
	if err := r.store.db.QueryRow(
		"UPDATE currency SET price=$2, volume=$3, last_trade=$4 WHERE symbol=$1 RETURNING  price, volume, last_trade",
		c.Symbol,
		c.Price,
		c.Volume,
		c.LastTrade,
	).Scan(
		&c.Price,
		&c.Volume,
		&c.LastTrade,
	); err != nil {
		log.Print(err)
		return nil, err
	}

	return &c, nil
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

	cur_array := make(map[string]model.CurrencyOutput)

	for rows.Next() {
		var raw_cur model.Currency
		if err := rows.Scan(&raw_cur.ID, &raw_cur.Symbol, &raw_cur.Price, &raw_cur.Volume,
			&raw_cur.LastTrade); err != nil {
			return nil, err
		}

		cur_array[raw_cur.Symbol] = model.CurrencyOutput{
			Price:     raw_cur.Price,
			Volume:    raw_cur.Volume,
			LastTrade: raw_cur.LastTrade,
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cur_array, nil
}
