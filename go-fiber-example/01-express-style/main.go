package main

import "github.com/gofiber/fiber"

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)
	app.Post("/body", func(c *fiber.Ctx) {
		c.Body()
		c.Send(c.Body())
	})
	// Create new sample GET routes
	app.Get("/demo", handler())
	app.Get("/list", handler())

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// start server
	app.Listen(3000)
}

// Handler
func hello(c *fiber.Ctx) {
	c.Send("Hello, world!")
}

// handler function
func handler() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.Send("This is a dummy route")
	}
}
