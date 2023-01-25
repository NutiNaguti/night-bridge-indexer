package main

import (
	"fmt"
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

	// Middleware
	e.Pre(apiVersion)

	// Routings version 1.0
	e.GET("v1/", controller.GetInfo)
	e.GET("v1/tx/last", controller.GetLastTransaction)
	e.GET("v1/tx", controller.GetTransactionsFromTo)
	// Admin functions
	e.PUT("v1/create", controller.AddTransaction)

	// Set db connection string
	db.SetupConnectionString(conf.Db.ConnectionString)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Header based versioning
func apiVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header
		apiVer := headers.Get("version")
		req.URL.Path = fmt.Sprintf("/%s%s", apiVer, req.URL.Path)
		return next(c)
	}
}
