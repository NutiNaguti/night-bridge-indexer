package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetInfo(c echo.Context) error {
	return c.String(http.StatusOK, "This is test controller for Night Bridge Indexer")
}
