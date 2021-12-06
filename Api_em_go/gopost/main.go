package main

import "github.com/gofiber/fiber"

type User struct {
	Name string
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello, Fiber!")
	})

	app.Get("/:name", func(ctx *fiber.Ctx) {
		name := ctx.Params("name")
		ctx.Send("My Name is: " + name)
	})

	app.Post("/criar", func(ctx *fiber.Ctx) {
		var user = &User{}

		ctx.BodyParser(user)

		ctx.JSON(user)
	})

	app.Listen("3001")
}
