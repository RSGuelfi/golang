package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Post("/method", func(ctx *fiber.Ctx) {
		if ctx.Method() == "POST" {
			ctx.Send("Via POST")
		}
	})

	app.Listen("3001")
}
