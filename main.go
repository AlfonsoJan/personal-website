package main

import (
	"log"
	"os"

	"github.com/alfonsojan/personal-website/setup"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := setup.Setup(e); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
