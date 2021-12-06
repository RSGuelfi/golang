package main

import (
	"github.com/RSGuelfi/golang/bancoDeDados/models"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB // Conn armazena as funções do gorm

func init() {
	var err error

	conn, err = gorm.Open("sqlite3", "product.db")

	if err != nil {
		panic("Houve um erro na tentativa de se conectar ao banco de dados.")
	}

	println("Banco de dados está online!")
}

func main() {
	app := fiber.New()

	var products []models.Product

	conn.AutoMigrate(&models.Product{})

	conn.Create(&models.Product{
		Name:  "XBOX",
		Price: 200.99,
		Stock: false,
	})

	conn.Create(&models.Product{
		Name:  "123+numberty",
		Price: 0.00,
		Stock: true,
	})

	// conn.Create(&models.Product{
	// 	Name:  "Playstation",
	// 	Price: 450.90,
	// 	Stock: t rue,
	// })

	conn.Delete(&models.Product{}, []int{8, 9}) // exclui um/multiplos pelo valor do id, Conn.Unscoped().Delete() = deleta dados permanentemente

	app.Get("/", func(ctx *fiber.Ctx) {
		conn.Find(&products)
		ctx.JSON(products)
	})
	app.Get("/:name", func(ctx *fiber.Ctx) {
		var product models.Product
		conn.Where("name = ?", ctx.Params("name")).Find(&product)

		ctx.JSON(product)
	})

	app.Get("/:id", func(ctx *fiber.Ctx) {
		var product models.Product
		conn.Where("id = ?", ctx.Params("id")).First(&product)

		ctx.JSON(product)
	})

	app.Listen(3001)
}
