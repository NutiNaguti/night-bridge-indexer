package main

import (
	"fmt"

	"github.com/NutiNaguti/night-bridge-indexer/common"
	"github.com/NutiNaguti/night-bridge-indexer/db"
	"github.com/NutiNaguti/night-bridge-indexer/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configs
	conf := common.New()

	e := echo.New()

	// Middleware
	e.Pre(apiVersion)
	// e.Use(middleware.BasicAuth())
	// e.Use(middleware.CORS())

	// Routings version 1.0
	e.GET("/", handler.GetInfo)
	e.GET("v1/tx/last", handler.GetLastTransaction)
	e.GET("v1/tx", handler.GetTransactionsFromTo)
	// Admin functions
	e.PUT("v1/create", handler.AddTransaction)

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
