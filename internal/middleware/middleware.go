package middleware

import (
	"strings"
	"time"

	"github.com/alfonsojan/personal-website/internal/utils/logger"
	"github.com/labstack/echo/v4"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		if strings.HasSuffix(c.Request().URL.String(), ".css") ||
			strings.HasSuffix(c.Request().URL.String(), ".js") ||
			strings.HasSuffix(c.Request().URL.String(), ".png") ||
			strings.HasSuffix(c.Request().URL.String(), ".ico") ||
			strings.HasSuffix(c.Request().URL.String(), ".jpg") {
			return nil
		}

		start := time.Now()

		duration := time.Since(start)

		logger.Logger.Info(
			c.Request().Method,
			c.Request().URL.String(),
			c.Response().Status,
			duration,
		)

		return nil
	}
}
