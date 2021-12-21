package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Motorista struct {
	gorm.Model
	Name      string  `json:"name"`
	Cpf       uint    `json:"cpf"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Available bool    `json:"available"`
	Category  string  `json:"category"`
}

type Passageiro struct {
	gorm.Model
	Name      string  `json:"name"`
	Cpf       uint    `json:"cpf"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (motorista Motorista) MotToString() string {
	return fmt.Sprintf("id: %d\nname: %s", motorista.ID, motorista.Name)
}

func (passageiro Passageiro) PasToString() string {
	return fmt.Sprintf("id: %d\nname: %s", passageiro.ID, passageiro.Name)
}

func IndexMot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var motorista Motorista
	db.First(&motorista, params["id"])
	if motorista.ID == 0 {
		json.NewEncoder(w).Encode("Motorista does not exists.")
	} else {
		json.NewEncoder(w).Encode(motorista)
	}
}

func ListMot(w http.ResponseWriter, r *http.Request) {
	var motorista []Motorista
	db.Find(&motorista)
	json.NewDecoder(r.Body).Decode(&motorista)
	json.NewEncoder(w).Encode(motorista)
}

func CreateMot(w http.ResponseWriter, r *http.Request) {
	var motorista Motorista
	json.NewDecoder(r.Body).Decode(&motorista)
	db.Create(&motorista)
	json.NewEncoder(w).Encode(motorista)
}

func PatchMot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var motorista Motorista
	db.First(&motorista, params["id"])
	json.NewDecoder(r.Body).Decode(&motorista)
	db.Save(&motorista)
	if motorista.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(motorista)
	json.NewEncoder(w).Encode("Motorista Updated.")
}

func UpdateMot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var motorista Motorista
	db.First(&motorista, params["id"])
	json.NewDecoder(r.Body).Decode(&motorista)
	db.Save(&motorista)
	if motorista.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(motorista)
	json.NewEncoder(w).Encode("Motorista Updated.")
}

func DeleteMot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var motorista Motorista

	error := db.Delete(&motorista, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("Motorista does not exists.")
	} else {
		json.NewEncoder(w).Encode("Motorista Deleted.")
	}
}

func IndexPas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var passageiro Passageiro
	db.First(&passageiro, params["id"])
	if passageiro.ID == 0 {
		json.NewEncoder(w).Encode("Passageiro does not exists.")
	} else {
		json.NewEncoder(w).Encode(passageiro)
	}
}

func ListPas(w http.ResponseWriter, r *http.Request) {
	var passageiro []Passageiro
	db.Find(&passageiro)
	json.NewDecoder(r.Body).Decode(&passageiro)
	json.NewEncoder(w).Encode(passageiro)
}

func CreatePas(w http.ResponseWriter, r *http.Request) {
	var passageiro Passageiro
	json.NewDecoder(r.Body).Decode(&passageiro)
	db.Create(&passageiro)
	// if id != 0 {
	// id + 1
	// }
	json.NewEncoder(w).Encode(passageiro)
}

func PatchPas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var passageiro Passageiro
	db.First(&passageiro, params["id"])
	json.NewDecoder(r.Body).Decode(&passageiro)
	db.Save(&passageiro)
	if passageiro.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(passageiro)
	json.NewEncoder(w).Encode("Passageiro Updated.")
}

func UpdatePas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var passageiro Passageiro
	db.First(&passageiro, params["id"])
	json.NewDecoder(r.Body).Decode(&passageiro)
	db.Save(&passageiro)
	if passageiro.ID == 0 {
		json.NewEncoder(w).Encode("Invalid ID.")
	}
	json.NewEncoder(w).Encode(passageiro)
	json.NewEncoder(w).Encode("Passageiro Updated.")
}

func DeletePas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var passageiro Passageiro

	error := db.Unscoped().Delete(&passageiro, params["id"])
	db.First(&passageiro, params["id"])

	fmt.Print("error", error)

	if error == nil {
		json.NewEncoder(w).Encode("Passageiro does not exists.")
	} else if !passageiro.DeletedAt.Valid {
		json.NewEncoder(w).Encode("Passageiro Already Deleted.")
		json.NewEncoder(w).Encode(passageiro.DeletedAt.Valid == error.NowFunc().IsZero())
	} else if error != nil {
		json.NewEncoder(w).Encode("Passageiro Deleted.")
		// if db.First(&user, 77).RecordNotFound() {
		// 	log.Println("UUID:", 77, err)
		// }
	}
	// if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	//   }

	//   models.DB.Delete(&book)

	//   c.JSON(http.StatusOK, gin.H{"data": true})
}

// func IndexPas(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var passageiro Passageiro
// 	db.First(&passageiro, params["id"])
// 	if passageiro.ID == 0 {
// 		json.NewEncoder(w).Encode("Passageiro does not exists.")
// 	} else {
// 		json.NewEncoder(w).Encode(passageiro)
// 	}
// }// Match

var dirValue float64

const radius = 6371 // Earth radius in kilometers

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (origin Motorista) Distance(destination Passageiro) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	dirValue = radius * c

	return dirValue
}

func Calcular(a float64, b float64) float64 {
	return a * b
}

func findMotorista(passageiro Passageiro) {
	var motoristas []Motorista
	km := 2.45
	min := motoristas[0]
	for i, motorista := range motoristas {
		if motorista.Available && motorista.Distance(passageiro) <= min.Distance(passageiro) {
			min = motorista
			tes := min.Distance(passageiro)
			defer fmt.Printf("\nMotorista %v: %s %.3f Km", i, min.Name, tes)
			for km = 2.45; 20 < i; i++ {
				fmt.Printf("%d %v %v", i, km, tes*2.4)
			}
		}
	}
}

// esse endpoint deve achar o motorista mais adequado para essa corrida seguindo os criterios
// deve estar com available == true
// deve ter a category compativel com a da requisicao podendo ser uber-x, uber-black, uber-pool
// deve ser o motorista mais proximos em km do location do passageiro
// calcular o valor no retorno da resposta http, cada km vai ser equivalente a 2,45 reais.
// post      /v1/match // enviar o payload { id_passageiro, category, location: { latitude, longitude }, destination: { latitude, longitude } }

func Match(w http.ResponseWriter, r *http.Request) {
	var motorista Motorista
	var passageiro Passageiro
	params := mux.Vars(r)
	db.First(&motorista, &passageiro, params["id"])
	if motorista.ID == 0 {
		json.NewEncoder(w).Encode("Motorista does not exists.")
	} else {
		findMotorista(passageiro)
		json.NewEncoder(w).Encode(motorista)
		json.NewEncoder(w).Encode(passageiro)
	}
	// MotToString()
	// pegar passageiro.id e usar a função distance
	// com a latitude e longitude do passageiro
	// e do motorista
	// criar um novo request entre motorista e passageiro
}

func Routes() {
	r := mux.NewRouter()

	// Motorista:
	r.HandleFunc("/motoristas/{id}", IndexMot).Methods("GET")
	r.HandleFunc("/motoristas", ListMot).Methods("GET")
	r.HandleFunc("/motoristas", CreateMot).Methods("POST")
	r.HandleFunc("/motoristas/{id}", PatchMot).Methods("PATCH")
	r.HandleFunc("/motoristas/{id}", UpdateMot).Methods("PUT")
	r.HandleFunc("/motoristas/{id}", DeleteMot).Methods("DELETE")

	// Passageiro:
	r.HandleFunc("/passageiros/{id}", IndexPas).Methods("GET")
	r.HandleFunc("/passageiros", ListPas).Methods("GET")
	r.HandleFunc("/passageiros", CreatePas).Methods("POST")
	r.HandleFunc("/passageiros/{id}", PatchPas).Methods("PATCH")
	r.HandleFunc("/passageiros/{id}", UpdatePas).Methods("PUT")
	r.HandleFunc("/passageiros/{id}", DeletePas).Methods("DELETE")

	// Match:
	r.HandleFunc("/match", Match).Methods("POST")

	log.Fatal(http.ListenAndServe(":4123", r))
}

func initDatabase() {

	var err error
	// db, err = gorm.Open(sqlite.Open("uber_api.db"), &gorm.Config{})
	dsn := "root:root@tcp(127.0.0.1:3306)/uber_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Database Connected...")

	db.AutoMigrate(&Motorista{}, &Passageiro{})
	fmt.Println("Database Migrated")
}

func main() {
	Motorista.MotToString(Motorista{})
	Passageiro.PasToString(Passageiro{})
	app := fiber.New()
	initDatabase()
	app.Use(logger.New())
	Routes()
}
