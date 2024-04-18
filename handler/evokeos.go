//go:build evokeos

package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "evokeos, config path: "+h.Config.Path)
}

func (h Handler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "evoke creative, config path: "+h.Config.Path)
}
