package setup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) error {
	return startServerWithGracefulShutdown(e)
}

func startServerWithGracefulShutdown(e *echo.Echo) error {
	errChan := make(chan error, 1)
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("error during startup: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")

	select {
	case err := <-errChan:
		return err
	default:
	}

	return nil
}
