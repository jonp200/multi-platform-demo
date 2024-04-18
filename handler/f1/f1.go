package f1Handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/os/configuration"
)

type F1Handler struct {
	Config configuration.Configuration
}

func (f F1Handler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "f1, config path: "+f.Config.Path)
}

func (f F1Handler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "flowbird, config path: "+f.Config.Path)
}
