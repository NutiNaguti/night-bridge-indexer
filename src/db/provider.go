package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var connectionSrting string

func CreateConnection() (*pgx.Conn, error) {
	var conn *pgx.Conn
	if connectionSrting == "" {
		return conn, errors.New("Connection string was empty")
	}

	conn, err := pgx.Connect(context.TODO(), connectionSrting)
	if err != nil {
		return conn, err
	}
	return conn, err
}

func SetupConnectionString(connString string) {
	connectionSrting = connString
}
