package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorRouting(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	if code == http.StatusNotFound {
		c.JSON(http.StatusNotFound, &errorResponse{"error", http.StatusNotFound, "The route you requested was not found"})
		return
	}
	c.JSON(http.StatusNotFound, &errorResponse{"error", http.StatusInternalServerError, "An unexpected error occurred"})
}
