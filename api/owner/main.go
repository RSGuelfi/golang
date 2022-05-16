package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name   string  `json:"name"`
	Modelo string  `json:"model"`
	Brand  string  `json:"brand"`
	Tipo   string  `json:"type"`
	Value  float64 `json:"value"`
	Owners []Owner `gorm:"many2many:owner_cars;not null;auto_preload" json:"owners"`
}

type Owner struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Cars    []Car  `gorm:"many2many:owner_cars;not null;auto_preload" json:"cars"`
}

var db *gorm.DB

// Formato JSON Automatico: w.Header().Set("Content-Type", "application/json")

func Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	OwnerID, err := params["OwnerID"]
	var car Car
	//var owner Owner
	if err {
		json.NewEncoder(w).Encode("OwnerID is required")
	}

	// buscar owner pelo id

	// ver como criar many to many no gorm

	// se o usuario não existir rejeitar a transferencia de carro para o owner

	fmt.Println(OwnerID)

	json.NewDecoder(r.Body).Decode(&car)
	db.Create(&car).Association("Owners").Append(&Owner{Cars: nil})
	db.Model(&car).Association("owner_cars").Append(&car)
	json.NewEncoder(w).Encode(car)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car []Car
	db.Find(&car)
	json.NewEncoder(w).Encode(car)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
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
	db.First(&car, params["id"])
	json.NewDecoder(r.Body).Decode(&car)
	db.Save(&car)
	if car.ID == 0 {
		json.NewEncoder(w).Encode("0 is not a valid id")
	}
	json.NewEncoder(w).Encode(car)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car

	error := db.Delete(&car, params["id"])

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

	db.Create(&owner)

	json.NewEncoder(w).Encode(owner)
}

func ListOwners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var owner []Owner
	db.Find(&owner)
	json.NewDecoder(r.Body).Decode(&owner)
	json.NewEncoder(w).Encode(owner)
}

func IDOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var owner Owner
	db.First(&owner, params["id"])
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
	db.First(&owner, params["id"])
	json.NewDecoder(r.Body).Decode(&owner)
	db.Save(&owner)
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

	error := db.Delete(&owner, params["id"])

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
	db, err = gorm.Open(sqlite.Open("owners.db"), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	db.AutoMigrate(&Car{}, &Owner{})
	fmt.Println("Database Migrated")
}

func (car Car) CarToString() string {
	return fmt.Sprintf("id: %d\nname: %s", car.ID, car.Name)
}

func (owner Owner) OwnerToString() string {
	return fmt.Sprintf("id: %d\nname: %s", owner.ID, owner.Name)
}
func main() {
	Car.CarToString(Car{})
	Owner.OwnerToString(Owner{})
	app := fiber.New()
	initDatabase()
	app.Use(logger.New())
	Routes()
}
