package defaultHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/os/configuration"
)

type DefaultHandler struct {
	Config configuration.Configuration
}

func (d DefaultHandler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "not implemented, config path: "+d.Config.Path)
}

func (d DefaultHandler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "not implemented, config path: "+d.Config.Path)
}
