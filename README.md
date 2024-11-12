# Real-Time Chat App with Go, Fiber, and HTMX

This project demonstrates how to build a simple real-time chat application using Go, the Fiber framework, and HTMX.

## Prerequisites

- Go 1.22 or higher
- Basic understanding of Go and HTTP servers

## Table of Contents

- [Getting Started](#getting-started)
- [Installing Dependencies](#installing-dependencies)
- [Setting Up the Project](#setting-up-the-project)
- [WebSocket Implementation](#websocket-implementation)
- [HTMX Integration](#htmx-integration)
- [Conclusion](#conclusion)

## Getting Started

1. Create a new directory for your project.
2. Initialize the Go module:  
   `go mod init github.com/yourusername/go-chat`

## Installing Dependencies

Install the required libraries by running:

```bash
go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/websocket/v2
go get -u github.com/gofiber/template/html/v2
```

## Setting Up the Project
* Create necessary folders for static files (static) and templates (views).
* Set up the main.go file to initialize a Fiber web server.
* Configure static file routes and views engine.

## WebSocket Implementation
* Create a websocket.go file to manage WebSocket connections.
* Implement methods to handle new clients, broadcast messages, and keep connections alive.

## HTMX Integration
* Add HTMX and WebSocket support to your frontend HTML files.
* Use HTMX's WebSocket extension to establish real-time communication between the client and server.

## Conclusion
You now have a simple real-time chat application built with Go, Fiber, and HTMX! You can extend this project by adding features like authentication, user management, and client IDs.

For more information, visit the [FreeCodeCamp - Real time chat with go, fiber and htmx](https://www.freecodecamp.org/news/real-time-chat-with-go-fiber-htmx/)
