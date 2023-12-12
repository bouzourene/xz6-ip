package main

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/bouzourene/xz6-ip/controllers"
	"github.com/bouzourene/xz6-ip/middlewares"
	"github.com/bouzourene/xz6-ip/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/jet/v2"
)

//go:embed static/*
var embedStatic embed.FS

//go:embed views
var embedViews embed.FS

func main() {
	// Very first thing we do is start the logger
	logger := tools.GetLogger()

	// Load config on application start
	logger.Info("Loading config from .env file")
	tools.LoadConfig()

	// Create templates engine
	engine := jet.NewFileSystem(
		http.FS(embedViews),
		".jet",
	)

	// Create webapp
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add global vars to template engine
	engine.Templates.AddGlobal("copyright_year", time.Now().Year())

	// Serve static content
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedStatic),
		PathPrefix: "static",
		Browse:     false,
	}))

	// Apply middlewares
	app.Use(middlewares.DefaultHeaders)

	// Define routes
	app.Get("/", controllers.Index)
	app.Get("/altproto", controllers.AltProto)
	app.Get("/json", controllers.ApiJson)
	app.Get("/xml", controllers.ApiXml)
	app.Get("/ip", controllers.ApiTextIp)
	app.Get("/reverse", controllers.ApiTextReverse)
	app.Get("/version", controllers.ApiTextVersion)
	app.Get("/bing", controllers.BingProxy)

	bindStr := fmt.Sprintf(
		"%s:%s",
		tools.ConfigValue("BIND_ADDR"),
		tools.ConfigValue("BIND_PORT"),
	)

	logger.Info(fmt.Sprintf(
		"Webserver starting on address: %s",
		bindStr,
	))

	err := app.Listen(bindStr)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
