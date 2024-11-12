package main

import (
	"go-chat/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {

	// Crrate views engine
	viewsEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Static route and directory
	app.Static("/static/", "./static")

	appHandler := handlers.NewAppHandler()
	server := NewWebSocket()

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	app.Get("/", appHandler.HandleGetIndex)

	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	// Start the http server
	app.Listen(":3000")
}
