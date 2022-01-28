package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
}

// 						user

func IndexUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	if user.ID == 0 {
		json.NewEncoder(w).Encode("user does not exists.")
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	var user []User
	db.Find(&user)
	json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	if user.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode("user Updated.")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User

	error := db.Delete(&user, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("user does not exists.")
	} else {
		json.NewEncoder(w).Encode("user Deleted.")
	}
}

func Routes() {
	r := mux.NewRouter()

	// user:
	r.HandleFunc("/user/{id}", IndexUser).Methods("GET")
	r.HandleFunc("/user", ListUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", PatchUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8888", r))
}

func initDatabase() {

	var err error
	// db, err = gorm.Open("postgres", "host=localhost port=4321 user=postgres dbname=your_db password=your_pass")
	// db, err = gorm.Open(sqlite.Open("uber_api.db"), &gorm.Config{})
	dsn := "host=localhost user=postgres password=postgres dbname=golangdocker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	db.AutoMigrate(&User{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Use(logger.New())
	Routes()
}
