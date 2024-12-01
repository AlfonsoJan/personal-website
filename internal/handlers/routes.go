package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "TEST!")
	})

	e.HTTPErrorHandler = ErrorRouting
}
