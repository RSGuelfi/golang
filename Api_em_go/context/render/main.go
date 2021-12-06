package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(&fiber.Settings{
		Views: html.New("./templates", ".html"),
	})
	app.Get("/", func(c *fiber.Ctx) {
		c.Render("person", fiber.Map{
			"name": "Arduino",
			"age":  39,
		})
	})
	app.Listen(3001)
}
