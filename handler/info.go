package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetInfo(c echo.Context) error {
	return c.String(http.StatusOK, `
	This is info controller for Night Bridge Indexer.
	This API is ver 1.0, you should add in requests header "version: v1"
Help:
	/tx/last - get last indexed transacion
	/tx?timestamp=...&page_token=...&page_size=... - get all transactions processed from timestamp since page equals page_token and with amount equals page_size
	/tx?from=...&to=... - get all transactions processed between two addresses`)
}
