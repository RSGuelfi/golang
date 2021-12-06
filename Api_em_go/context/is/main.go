package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/is", func(ctx *fiber.Ctx) {
		if ctx.Is("json") {
			ctx.Send("JSON enviando com sucesso")
		} else {
			ctx.Status(400).Send("O conteúdo deve estar no formato JSON")
		}
	})

	app.Listen("3001")
}
