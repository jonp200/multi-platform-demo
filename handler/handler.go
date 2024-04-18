package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	GetOSName(c echo.Context) error
	GetOrg(c echo.Context) error
}
