package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// Menambahkan middleware logger
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} ${method} ${path} ${latency}\n",
	}))

	// Menambahkan CORS middleware untuk mengizinkan semua permintaan lintas asal
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Mengizinkan semua domain, sesuaikan sesuai kebutuhan
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup proxy routes
	app.Get("/books", proxy.Forward("http://localhost:8080/books"))
	app.Post("/books", proxy.Forward("http://localhost:8080/books"))
	app.Put("/books/:books_id", proxy.Forward("http://localhost:8080/books/:books_id")) // Corrected to use books_id if that's what you are using in your service
	app.Delete("/books/:books_id", proxy.Forward("http://localhost:8080/books/:books_id"))

	app.Listen(":8082") // Changed from 8080 to 8082 to avoid port conflict
}
