package main

import (
	"log"

	"github.com/NutiNaguti/night-bridge-indexer/common"
	"github.com/NutiNaguti/night-bridge-indexer/controller"
	"github.com/NutiNaguti/night-bridge-indexer/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file")
	}
}

func main() {
	// Load configs
	conf := common.New()

	e := echo.New()

	// Routings
	e.GET("/", controller.GetInfo)
	e.GET("/tx", controller.GetLastTransaction)
	e.GET("/tx", controller.GetTransactionsFromTo)
	e.PUT("/create", controller.AddTransaction)

	// Set db connection string
	db.SetupConnectionString(conf.Db.ConnectionString)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
