package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// Setup proxy routes
	app.Get("/books", proxy.Forward("http://localhost:8080/books"))
	app.Post("/books", proxy.Forward("http://localhost:8080/books"))
	app.Put("/books/:books_id", proxy.Forward("http://localhost:8080/books/:books_id")) // Corrected to use books_id if that's what you are using in your service
	app.Delete("/books/:books_id", proxy.Forward("http://localhost:8080/books/:books_id"))

	app.Listen(":8082") // Changed from 8080 to 8082 to avoid port conflict
}
