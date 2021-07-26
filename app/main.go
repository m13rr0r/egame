package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/m13rr0r/egame/handler"
	"github.com/m13rr0r/egame/infrastructure/storage"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Logger.Print(os.Getenv("APP_DB_STORAGE_URI"))
	stg, err := storage.NewStorage(os.Getenv("APP_DB_STORAGE_URI"))
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler.Handler{DB: stg}
	// Routes
	e.POST("/event", h.CreateEvent)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}



