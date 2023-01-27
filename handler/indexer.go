package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NutiNaguti/night-bridge-indexer/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetLastTransaction(c echo.Context) error {
	var tx *model.Transaction
	tx, err := model.GetLastTransaction(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, tx)
}

func GetTransactionsFromTo(c echo.Context) error {
	var txs model.TransactionsPage
	var err error
	timestamp := c.QueryParam("timestamp")
	pageToken, err := strconv.ParseUint(c.QueryParam("page_token"), 10, 16)
	if err != nil {
		return err
	}
	pageSize, err := strconv.ParseUint(c.QueryParam("page_size"), 10, 16)
	if err != nil {
		return err
	}
	txs, err = model.GetTransactionsSince(context.Background(), timestamp, uint16(pageToken), uint16(pageSize))
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, txs)
}

func AddTransaction(c echo.Context) error {
	sender := c.QueryParam("sender")
	receiver := c.QueryParam("receiver")
	amount := c.QueryParam("amount")
	timestamp := c.QueryParam("timestamp")
	err := model.CreateTransaction(context.Background(), sender, receiver, amount, timestamp)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "")
}
