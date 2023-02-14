package model

import (
	"context"

	"github.com/NutiNaguti/night-bridge-indexer/db"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/shopspring/decimal"
)

const maxPageSize uint16 = 100

type TransactionsPage struct {
	Txs           []Transaction
	NextPageToken uint16
}

type Transaction struct {
	Id        int             `json:"id" db:"id"`
	Sender    string          `json:"from" db:"sender"`
	Receiver  string          `json:"to" db:"receiver"`
	Amount    decimal.Decimal `json:"amount" db:"amount"`
	Timestamp int64           `json:"timestamp" db:"timestamp"`
}

func GetLastTransaction(ctx context.Context, out chan Transaction) error {
	var tx Transaction
	conn, err := db.CreateConnection()
	if err != nil {
		log.Error(err)
		return err
	}

	err = conn.QueryRow(ctx, `select * from transactions order by timestamp desc limit 1`).Scan(&tx.Id, &tx.Sender, &tx.Receiver, &tx.Amount, &tx.Timestamp)
	conn.Close(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	out <- tx
	return err
}

func GetTransactionsSince(ctx context.Context, timestamp string, pageToken uint16, pageSize uint16, out chan TransactionsPage) error {
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	var txs []Transaction
	var txsResponse TransactionsPage

	conn, err := db.CreateConnection()
	if err != nil {
		return err
	}

	rows, _ := conn.Query(ctx, `select * from transactions where timestamp >= $1 order by timestamp limit $2 offset $3`, timestamp, pageSize, pageSize*(pageToken-1))
	conn.Close(ctx)

	if err != nil {
		return err
	}

	txs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[Transaction])
	if err != nil {
		return err
	}

	txsResponse.Txs = txs
	txsResponse.NextPageToken = pageToken + 1

	out <- txsResponse
	return err
}

func CreateTransaction(ctx context.Context, sender string, receiver string, amount string, timestamp string, out chan bool) error {
	conn, err := db.CreateConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `insert into transactions (sender, receiver, amount, timestamp) values ($1, $2, $3, $4)`, sender, receiver, amount, timestamp)
	conn.Close(ctx)

	if err != nil {
		return err
	}

	out <- true
	return nil
}
