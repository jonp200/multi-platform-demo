package main

import (
	"github.com/labstack/echo/v4"
	"multi-platform-demo/handler"
	"multi-platform-demo/os/configuration"
)

func main() {
	e := echo.New()
	c := configuration.Configuration{Port: "8080"}

	h := handler.Handler{Config: c}

	e.GET("/os", h.GetOSName)
	e.GET("/org", h.GetOrg)

	e.Logger.Fatal(e.Start(":" + c.Port))
}
