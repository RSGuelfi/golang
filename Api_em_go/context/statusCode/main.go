package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusInternalServerError).Send("Hello World!")
	})

	app.Listen("3001")
}
