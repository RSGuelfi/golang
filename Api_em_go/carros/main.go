package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Car struct {
	gorm.Model
	Name   string  `json:"name"`
	Modelo string  `json:"model"`
	Brand  string  `json:"brand"`
	Tipo   string  `json:"type"`
	Value  float64 `json:"value"`
}

var DB *gorm.DB

// func Index(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	var car Car
// 	db.Find(&car, id)
// 	c.JSON(car)
// }

// func List(c *fiber.Ctx) {
// 	var cars []Car
// 	db.Find(&cars)
// 	c.JSON(cars)
// }

// func Store(c *fiber.Ctx) {
// 	car := new(Car)
// 	if err := c.BodyParser(car); err != nil {
// 		c.Status(503).Send(err)
// 		return
// 	}

// 	db.NewRecord(&car)
// 	db.Create(&car)
// 	c.Status(201)
// 	c.JSON(car)
// }

// func Update(c *fiber.Ctx) {
// 	id := c.Params("id")
// 	var car Car
// 	db.First(&car, id)

// 	db.NewRecord(&car)
// 	db.Model(&car).Updates(map[string]interface{}{
// 		"name":   car.Name,
// 		"modelo": car.Modelo,
// 		"brand":  car.Brand,
// 		"tipo":   car.Tipo,
// 		"value":  car.Value,
// 	})
// 	car.Name = Car{}.Name
// 	car.Modelo = Car{}.Modelo
// 	car.Brand = Car{}.Brand
// 	car.Tipo = Car{}.Tipo
// 	car.Value = Car{}.Value
// 	db.AutoMigrate(&car)
// 	db.Save(&car)
// 	c.JSON(car)
// }
//
// func Delete(w http.ResponseWriter, r *http.Request) {
// 	id := c.Params("id")

// 	var car Car
// 	db.First(&car, id)
// 	if car.Name == "" {
// 		c.Status(500).Send("Car not found")
// 		return
// 	}
// 	db.Delete(&car)
// 	c.Send("Car deleted")
// }

func Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	json.NewDecoder(r.Body).Decode(&car)
	DB.Create(&car)
	json.NewEncoder(w).Encode(car)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car []Car
	DB.Find(&car)
	json.NewEncoder(w).Encode(car)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	DB.First(&car, params["id"])
	if car.ID == 0 {
		json.NewEncoder(w).Encode("Car does not exists")
	} else {
		json.NewEncoder(w).Encode(car)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	DB.First(&car, params["id"])
	json.NewDecoder(r.Body).Decode(&car)
	DB.Save(&car)
	if car.ID == 0 {
		json.NewEncoder(w).Encode("0 is not a valid id")
	}
	json.NewEncoder(w).Encode(car)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car

	error := DB.Delete(&car, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("Car does not exists")
	} else {
		json.NewEncoder(w).Encode("Car Deleted.")
	}
}

// func Routes(r *fiber.App) {
// 	r.Get("/cars/:id", Index)
// 	r.Get("/cars", List)
// 	r.Post("/cars", Store)
// 	r.Patch("/cars/:id", UpdateCar)
// 	r.Delete("/cars/:id", Delete)
// }

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/cars", List).Methods("GET")
	r.HandleFunc("/cars/{id}", Index).Methods("GET")
	r.HandleFunc("/cars", Store).Methods("POST")
	r.HandleFunc("/cars/{id}", Update).Methods("PATCH")
	r.HandleFunc("/cars/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5543", r))
}

func initDatabase() {

	var err error
	DB, err = gorm.Open("sqlite3", "cars.db")
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	DB.AutoMigrate(&Car{})
	fmt.Println("Database Migrated")

}

func main() {

	app := fiber.New()
	initDatabase()
	defer DB.Close()
	app.Use(middleware.Logger())
	Routes()
}
