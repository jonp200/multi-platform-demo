package evokeHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/os/configuration"
)

type EvokeHandler struct {
	Config configuration.Configuration
}

func (e EvokeHandler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "evokeos, config path: "+e.Config.Path)
}

func (e EvokeHandler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "evoke creative, config path: "+e.Config.Path)
}
