package setup

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/alfonsojan/personal-website/internal/config"
	"github.com/alfonsojan/personal-website/internal/handlers"
	request "github.com/alfonsojan/personal-website/internal/middleware"
	"github.com/alfonsojan/personal-website/internal/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo) error {
	if err := logger.New("app.log"); err != nil {
		return fmt.Errorf("could not create custom logger: %v", err)
	}
	if err := config.SetConfigFile(); err != nil {
		return err
	}
	setupMiddleware(e)
	setupRoutes(e)
	return startServerWithGracefulShutdown(e)
}

func setupRoutes(e *echo.Echo) {
	e.File("favicon.ico", "./static/images/favicon.ico")
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: false,
	}))
	handlers.SetupRoutes(e)
}

func setupMiddleware(e *echo.Echo) {
	e.Use(request.Log)
}

func startServerWithGracefulShutdown(e *echo.Echo) error {
	port := fmt.Sprintf(":%s", strconv.Itoa(config.AppConfig.Server.Port))
	errChan := make(chan error, 1)
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("error during startup: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %v", err)
	}

	logger.Logger.Info("Server gracefully stopped")

	select {
	case err := <-errChan:
		return err
	default:
	}

	return nil
}
