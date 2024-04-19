package handler

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"multi-platform-demo/os/configuration"
)

type Handler struct {
	Config configuration.Configuration
}

// StubContextRecorder creates an echo.Context with the given request body and also its response recorder.
// Only used in unit tests.
func StubContextRecorder(body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	// There's no point logging in unit tests
	e.Logger.SetLevel(log.OFF)
	// Method and path does not matter here
	req := httptest.NewRequest(http.MethodGet, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}
