package model

import "github.com/shopspring/decimal"

type Transaction struct {
	Id        int             `json:"id"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Amount    decimal.Decimal `json:"amount"`
	Timestamp int64           `json:"timestamp"`
}

// func GetLastTransaction() (Transaction, error) {

// }
