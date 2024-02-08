package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/evantyip/deep_dive_library/internal/handler"
)

func main() {
  app := echo.New()
  app.Use(middleware.Logger())

  libraryHandler := handler.LibraryHandler{}

  app.GET("/", libraryHandler.HandleShow)
  
  app.Start(":3000")
}
