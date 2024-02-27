package main

import (
	"github.com/labstack/echo/v4"
	"github.com/slingercode/go-echo-templ-plus-whatever/handlers"
)

func main() {
  app := echo.New()

  app.Static("/static", "static")

  // Routes
  app.GET("/", handlers.HandleIndex)

  app.Logger.Fatal(app.Start(":3000"))
}
