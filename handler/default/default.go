package defaultHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/os/configuration"
)

type Handler struct {
	Config configuration.Configuration
}

func (h Handler) GetOSName(c echo.Context) error {
	return c.String(http.StatusOK, "not implemented, config path: "+h.Config.Path)
}

func (h Handler) GetOrg(c echo.Context) error {
	return c.String(http.StatusOK, "not implemented, config path: "+h.Config.Path)
}
