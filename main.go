package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/handler"
	"multi-platform-demo/os/configuration"

	_default "multi-platform-demo/handler/default"
	evokeos "multi-platform-demo/handler/evokeos"
	f1 "multi-platform-demo/handler/f1"
)

func main() {
	e := echo.New()
	c := configuration.Configuration{Path: "config.yaml"}
	t := os.Getenv("TARGET_OS")

	var h handler.Handler

	switch t {
	case "evokeos":
		h = evokeos.Handler{Config: c}
	case "f1":
		h = f1.Handler{Config: c}
	default:
		h = _default.Handler{Config: c}
	}

	e.GET("/os", h.GetOSName)
	e.GET("/org", h.GetOrg)

	e.Logger.Fatal(e.Start(":8080"))
}
