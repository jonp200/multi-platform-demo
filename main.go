package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"multi-platform-demo/handler"
	"multi-platform-demo/os/configuration"

	defaultHandler "multi-platform-demo/handler/default"
	evokeHandler "multi-platform-demo/handler/evoke"
	f1Handler "multi-platform-demo/handler/f1"
)

func main() {
	e := echo.New()
	c := configuration.Configuration{Path: "config.yaml"}
	t := os.Getenv("TARGET_OS")

	var h handler.Handler

	switch t {
	case "evokeos":
		h = evokeHandler.EvokeHandler{Config: c}
	case "f1":
		h = f1Handler.F1Handler{Config: c}
	default:
		h = defaultHandler.DefaultHandler{Config: c}
	}

	e.GET("/os", h.GetOSName)
	e.GET("/org", h.GetOrg)

	e.Logger.Fatal(e.Start(":8080"))
}
