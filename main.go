package main

import (
	"fmt"
	"os"

	"github.com/alfonsojan/personal-website/internal/utils/logger"
	"github.com/alfonsojan/personal-website/setup"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := setup.Setup(e); err != nil {
		logger.Logger.Error(fmt.Errorf("error: %v", err))
		os.Exit(2)
	}
}
