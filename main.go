package main

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/liuximu/flashcard/router"
	"github.com/liuximu/flashcard/shared"
)

const confDir = "./conf"
const appConfFile = "app.json"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Logger.SetLevel(log.DEBUG)

	// Routes
	router.Route(e)

	if err := shared.LoadAppConf(filepath.Join(confDir, appConfFile), e.Logger); err != nil {
		panic(err)
	}

	// Start server
	e.Logger.Fatal(e.Start(shared.AppConfSingleton.Port))
}
