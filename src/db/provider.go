package db

import (
	"log"

	"github.com/jackc/pgx"
)

var connectionSrting string

func CreateConnection() (*pgx.Conn, error) {
	if connectionSrting == "" {
		log.Fatal("Connection string was empty")
	}
	connConfig, err := pgx.ParseConnectionString(connectionSrting)
	conn, err := pgx.Connect(connConfig)
	return conn, err
}

func SetupConnectionString(connString string) {
	connectionSrting = connString
}

func CreateTable() {
	conn, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	conn.QueryRow(`create table Transactions(
			id serial primary key,
			output text not null,
			input text not null,
			amount decimal not null,
			timestamp integer not null,
			check (timestamp > 0),
			check (timestamp > 0)
	)`)

	conn.Close()
}

func CreateTransaction() {
	conn, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	conn.QueryRow(``)
	conn.Close()
}
