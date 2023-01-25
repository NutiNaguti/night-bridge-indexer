package controller

import (
	"log"
	"net/http"

	"github.com/NutiNaguti/night-bridge-indexer/model"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

func GetLastTransaction(c echo.Context) error {
	amount, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		log.Fatal(err)
	}
	tx := &model.Transaction{
		Id:        0,
		From:      "nutinaguti.testnet",
		To:        "0x8CAB5E96E1ab09e8678a8ffC75b5D818e73D4707",
		Amount:    amount,
		Timestamp: 1,
	}
	return c.JSON(http.StatusOK, tx)
}

func GetTransactionsFromTo(c echo.Context) error {
	amount, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		log.Fatal(err)
	}
	// from := c.Param("from")
	// to := c.Param("to")
	txs := []model.Transaction{
		{
			Id:        0,
			From:      "nutinaguti.testnet",
			To:        "0x8CAB5E96E1ab09e8678a8ffC75b5D818e73D4707",
			Amount:    amount,
			Timestamp: 1,
		},
		{
			Id:        1,
			From:      "nutinaguti.testnet",
			To:        "0x8CAB5E96E1ab09e8678a8ffC75b5D818e73D4707",
			Amount:    amount,
			Timestamp: 2,
		},
	}
	return c.JSON(http.StatusOK, txs)
}

func AddTransaction(c echo.Context) error {
	sender := c.QueryParam("sender")
	receiver := c.QueryParam("receiver")
	amount := c.QueryParam("amount")
	timestamp := c.QueryParam("timestamp")
	log.Printf("%s, %s, %s, %s", sender, receiver, amount, timestamp)
	return c.JSON(http.StatusOK, "")
}
