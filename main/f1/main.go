//go:build f1

package main

import (
	"github.com/labstack/echo/v4"
	"multi-platform-demo/handler"
	"multi-platform-demo/os/configuration"

	f1 "multi-platform-demo/handler/f1"
)

func main() {
	e := echo.New()
	c := configuration.Configuration{Path: "config.yaml"}

	var h handler.Handler = f1.Handler{Config: c}

	e.GET("/os", h.GetOSName)
	e.GET("/org", h.GetOrg)

	e.Logger.Fatal(e.Start(":8080"))
}
