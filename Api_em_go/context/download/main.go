package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/file", func(ctx *fiber.Ctx) {
		ctx.Download("./Foto.png", "Foto.png")
	})

	app.Listen("3001")
}
