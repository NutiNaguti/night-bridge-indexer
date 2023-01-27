package main

import (
	"fmt"
	"os"

	"github.com/NutiNaguti/night-bridge-indexer/common"
	"github.com/NutiNaguti/night-bridge-indexer/db"
	"github.com/NutiNaguti/night-bridge-indexer/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	// Load configs
	conf := common.New()

	// Create server
	e := echo.New()

	// Middleware
	e.Pre(apiVersion)
	// TODO: replace default config by custom
	e.Use(middleware.CORS())
	// Setup logging
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().Str("URI", v.URI).Int("status", v.Status).Msg("request")
			return nil
		},
	},
	))

	e.GET("/", handler.GetInfo)

	// Routings version 1.0
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
		if apiVer != "" {
			req.URL.Path = fmt.Sprintf("/%s%s", apiVer, req.URL.Path)
		}
		return next(c)
	}
}
