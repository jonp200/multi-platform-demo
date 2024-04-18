//go:build evokeos

package main

import (
	"github.com/labstack/echo/v4"
	"multi-platform-demo/handler"
	"multi-platform-demo/os/configuration"

	evokeos "multi-platform-demo/handler/evokeos"
)

func main() {
	e := echo.New()
	c := configuration.Configuration{Path: "config.yaml"}

	var h handler.Handler = evokeos.Handler{Config: c}

	e.GET("/os", h.GetOSName)
	e.GET("/org", h.GetOrg)

	e.Logger.Fatal(e.Start(":8080"))
}
