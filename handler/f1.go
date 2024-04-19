//go:build f1

package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "f1, port: "+h.Config.Port)
}

func (h Handler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "flowbird, port: "+h.Config.Port)
}
