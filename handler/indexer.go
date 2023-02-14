package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/NutiNaguti/night-bridge-indexer/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/cases"
)

var defaultTimeDuration = time.Second * 10

func GetLastTransaction(c echo.Context) error {
	d := time.Now().Add(defaultTimeDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	tx := make(chan model.Transaction)
	go model.GetLastTransaction(ctx, tx)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-tx:
		return c.JSON(http.StatusOK, result)
	}
}

func GetTransactionsFromTo(c echo.Context) error {
	d := time.Now().Add(defaultTimeDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	timestamp := c.QueryParam("timestamp")
	pageToken, err := strconv.ParseUint(c.QueryParam("page_token"), 10, 16)
	if err != nil {
		return err
	}
	pageSize, err := strconv.ParseUint(c.QueryParam("page_size"), 10, 16)
	if err != nil {
		return err
	}

	txs := make(chan model.TransactionsPage)
	go model.GetTransactionsSince(ctx, timestamp, uint16(pageToken), uint16(pageSize), txs)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-txs:
		return c.JSON(http.StatusOK, result)
	}
}

func AddTransaction(c echo.Context) error {
	d := time.Now().Add(defaultTimeDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	sender := c.QueryParam("sender")
	receiver := c.QueryParam("receiver")
	amount := c.QueryParam("amount")
	timestamp := c.QueryParam("timestamp")

	success := make(chan bool)
	go model.CreateTransaction(ctx, sender, receiver, amount, timestamp, success)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-success:
		return c.JSON(http.StatusOK, result)
	}
}
