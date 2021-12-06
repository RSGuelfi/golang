package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Client struct {
	gorm.Model
	Name    string `json:"name"`
	Address Address
}

type Address struct {
	gorm.Model
	Name         string `json:"name"`
	ClientID     uint
	RestaurantID uint
}

type Restaurant struct {
	gorm.Model
	Name    string `json:"name"`
	Address Address
}

type Product struct {
	gorm.Model
	Name         string `json:"name"`
	RestaurantID uint   `json:"id_restaurant"`
	Restaurant   Restaurant
}

type Request struct {
	gorm.Model
	RestaurantID uint
}

// Esse endpoint é responsavel por fechar o pedido, vamos receber o id do client e recuperar o address dele, em seguida vamos enviar o pedido de entrega pro resturante referente ao id do product
func TRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client Client
	// var product Product
	var request Request
	json.NewDecoder(r.Body).Decode(&request)
	db.Preload("product").First(&client, params["id"]).Create(&request)
	json.NewEncoder(w).Encode(request)
}

// CLIENT

func IndexCli(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client Client
	db.First(&client, params["id"])
	if client.ID == 0 {
		json.NewEncoder(w).Encode("client does not exists.")
	} else {
		json.NewEncoder(w).Encode(client)
	}
}

func ListCli(w http.ResponseWriter, r *http.Request) {
	var client []Client
	db.Find(&client)
	json.NewDecoder(r.Body).Decode(&client)
	json.NewEncoder(w).Encode(client)
}

func CreateCli(w http.ResponseWriter, r *http.Request) {
	var client Client
	json.NewDecoder(r.Body).Decode(&client)
	db.Create(&client)
	json.NewEncoder(w).Encode(client)
}

func PatchCli(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client Client
	db.First(&client, params["id"])
	json.NewDecoder(r.Body).Decode(&client)
	db.Save(&client)
	if client.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(client)
	json.NewEncoder(w).Encode("client Updated.")
}

func DeleteCli(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var client Client

	error := db.Delete(&client, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("client does not exists.")
	} else {
		json.NewEncoder(w).Encode("client Deleted.")
	}
}

// ADDRESS

func IndexAddr(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var address Address
	db.First(&address, params["id"])
	if address.ID == 0 {
		json.NewEncoder(w).Encode("address does not exists.")
	} else {
		json.NewEncoder(w).Encode(address)
	}
}

func ListAddr(w http.ResponseWriter, r *http.Request) {
	var address []Address
	db.Find(&address)
	json.NewDecoder(r.Body).Decode(&address)
	json.NewEncoder(w).Encode(address)
}

func CreateAddr(w http.ResponseWriter, r *http.Request) {
	var address Address
	json.NewDecoder(r.Body).Decode(&address)
	db.Create(&address)
	json.NewEncoder(w).Encode(address)
}

func PatchAddr(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var address Address
	db.First(&address, params["id"])
	json.NewDecoder(r.Body).Decode(&address)
	db.Save(&address)
	if address.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(address)
	json.NewEncoder(w).Encode("address Updated.")
}

func DeleteAddr(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var address Address

	error := db.Delete(&address, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("address does not exists.")
	} else {
		json.NewEncoder(w).Encode("address Deleted.")
	}
}

// RESTAURANT

func IndexRest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant Restaurant
	db.First(&restaurant, params["id"])
	if restaurant.ID == 0 {
		json.NewEncoder(w).Encode("restaurant does not exists.")
	} else {
		json.NewEncoder(w).Encode(restaurant)
	}
}

func ListRest(w http.ResponseWriter, r *http.Request) {
	var restaurant []Restaurant
	db.Find(&restaurant)
	json.NewDecoder(r.Body).Decode(&restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

func CreateRest(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	json.NewDecoder(r.Body).Decode(&restaurant)
	db.Create(&restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

func PatchRest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant Restaurant
	db.First(&restaurant, params["id"])
	json.NewDecoder(r.Body).Decode(&restaurant)
	db.Save(&restaurant)
	if restaurant.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(restaurant)
	json.NewEncoder(w).Encode("restaurant Updated.")
}

func DeleteRest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var restaurant Restaurant

	error := db.Delete(&restaurant, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("restaurant does not exists.")
	} else {
		json.NewEncoder(w).Encode("restaurant Deleted.")
	}
}

// PRODUCT

func IndexProd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	db.First(&product, params["id"])
	if product.ID == 0 {
		json.NewEncoder(w).Encode("product does not exists.")
	} else {
		json.NewEncoder(w).Encode(product)
	}
}

func ListProd(w http.ResponseWriter, r *http.Request) {
	var product []Product
	db.Find(&product)
	json.NewDecoder(r.Body).Decode(&product)
	json.NewEncoder(w).Encode(product)
}

func CreateProd(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	db.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func PatchProd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	db.First(&product, params["id"])
	json.NewDecoder(r.Body).Decode(&product)
	db.Save(&product)
	if product.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(product)
	json.NewEncoder(w).Encode("product Updated.")
}

func DeleteProd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product

	error := db.Delete(&product, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("product does not exists.")
	} else {
		json.NewEncoder(w).Encode("product Deleted.")
	}
}

func Routes() {
	r := mux.NewRouter()

	// Client:
	r.HandleFunc("/clients/{id}", IndexCli).Methods("GET")
	r.HandleFunc("/clients", ListCli).Methods("GET")
	r.HandleFunc("/clients", CreateCli).Methods("POST")
	r.HandleFunc("/clients/{id}", PatchCli).Methods("PATCH")
	r.HandleFunc("/clients/{id}", DeleteCli).Methods("DELETE")

	// Address:
	r.HandleFunc("/addresses/{id}", IndexAddr).Methods("GET")
	r.HandleFunc("/addresses", ListAddr).Methods("GET")
	r.HandleFunc("/addresses", CreateAddr).Methods("POST")
	r.HandleFunc("/addresses/{id}", PatchAddr).Methods("PATCH")
	r.HandleFunc("/addresses/{id}", DeleteAddr).Methods("DELETE")

	// Restaurant:
	r.HandleFunc("/restaurants/{id}", IndexRest).Methods("GET")
	r.HandleFunc("/restaurants", ListRest).Methods("GET")
	r.HandleFunc("/restaurants", CreateRest).Methods("POST")
	r.HandleFunc("/restaurants/{id}", PatchRest).Methods("PATCH")
	r.HandleFunc("/restaurants/{id}", DeleteRest).Methods("DELETE")

	// Product:
	r.HandleFunc("/products/{id}", IndexProd).Methods("GET")
	r.HandleFunc("/products", ListProd).Methods("GET")
	r.HandleFunc("/products", CreateProd).Methods("POST")
	r.HandleFunc("/products/{id}", PatchProd).Methods("PATCH")
	r.HandleFunc("/products/{id}", DeleteProd).Methods("DELETE")

	// Request:
	r.HandleFunc("/requests/{id}", TRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":5123", r))
}

func initDatabase() {

	var err error
	// db, err = gorm.Open("postgres", "host=localhost port=4321 user=postgres dbname=your_db password=your_pass")
	// db, err = gorm.Open(sqlite.Open("uber_api.db"), &gorm.Config{})
	dsn := "root:root@tcp(127.0.0.1:3306)/api_ifood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	db.AutoMigrate(&Client{}, &Address{}, &Restaurant{}, Product{}, Request{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Use(logger.New())
	Routes()
}
