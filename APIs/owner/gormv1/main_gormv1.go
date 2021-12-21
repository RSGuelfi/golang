package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Car struct {
	OwnerID uint
	Owners  []Owner `gorm:"many2many:owner_cars;not null" json:"owners"`
	gorm.Model
	Name   string  `json:"name"`
	Modelo string  `json:"model"`
	Brand  string  `json:"brand"`
	Tipo   string  `json:"type"`
	Value  float64 `json:"value"`
}

type Owner struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Cars    []Car  `gorm:"many2many:owner_cars;not null" json:"cars"`
}

var DB *gorm.DB

func Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	OwnerID, err := params["OwnerID"]
	var car Car
	if err {
		json.NewEncoder(w).Encode("idOwner is required")
	}

	// buscar owner pelo id

	// ver como criar many to many no gorm

	// se o usuario não existir rejeitar a transferencia de carro para o owner

	fmt.Println(OwnerID)

	json.NewDecoder(r.Body).Decode(&car)
	DB.Model(&Car{}).Preload("Owners").Find(OwnerID).Create(&car)
	DB.Model(&Owner{}).Association("Cars").Append(&Car{})
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

func NewOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var owner Owner

	json.NewDecoder(r.Body).Decode(&owner)

	DB.Create(&owner)

	json.NewEncoder(w).Encode(owner)
}

func ListOwners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var owner []Owner
	DB.Find(&owner)
	json.NewDecoder(r.Body).Decode(&owner)
	json.NewEncoder(w).Encode(owner)
}

func IDOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var owner Owner
	DB.First(&owner, params["id"])
	if owner.ID == 0 {
		json.NewEncoder(w).Encode("Owner does not exists")
	} else {
		json.NewEncoder(w).Encode(owner)
	}
}

func UpdateOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var owner Owner
	DB.First(&owner, params["id"])
	json.NewDecoder(r.Body).Decode(&owner)
	DB.Save(&owner)
	if owner.ID == 0 {
		json.NewEncoder(w).Encode("iis invalid id")
	}
	json.NewEncoder(w).Encode(owner)
	json.NewEncoder(w).Encode("Owner Updated.")
}

func DeleteOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var owner Owner

	error := DB.Delete(&owner, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("Owner does not exists")
	} else {
		json.NewEncoder(w).Encode("Owner Deleted.")
	}
}

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/cars/{OwnerID}", Store).Methods("POST")
	r.HandleFunc("/cars", List).Methods("GET")
	r.HandleFunc("/cars/{id}", Index).Methods("GET")
	r.HandleFunc("/cars/{id}", Update).Methods("PATCH")
	r.HandleFunc("/cars/{id}", Delete).Methods("DELETE")
	r.HandleFunc("/owners", NewOwner).Methods("POST")
	r.HandleFunc("/owners", ListOwners).Methods("GET")
	r.HandleFunc("/owners/{id}", IDOwner).Methods("GET")
	r.HandleFunc("/owners/{id}", UpdateOwner).Methods("PATCH")
	r.HandleFunc("/owners/{id}", DeleteOwner).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5543", r))
}

func initDatabase() {

	var err error
	DB, err = gorm.Open("sqlite3", "owners.db")
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	DB.AutoMigrate(&Car{}, &Owner{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	defer DB.Close()
	app.Use(logger.New())
	Routes()
}
