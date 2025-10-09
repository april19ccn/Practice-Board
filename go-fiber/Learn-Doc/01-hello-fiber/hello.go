package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// 静态文件，注意顺序，如果放在后面，会导致 app.Get("/" ... 先处理路由
	app.Static("/", "./public")

	// Simple route 简单路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Parameters 参数
	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// Optional parameter 可选参数
	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	// Wildcards 通配符
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => API path: user/john
	})

	app.Listen(":3000")
}
