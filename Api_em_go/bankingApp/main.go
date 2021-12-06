package main

import (
	"github.com/RSGuelfi/golang/bankingApp/api"
	"github.com/RSGuelfi/golang/bankingApp/database"
)

func main() {
	database.InitDatabase()

	// Do migration
	// migrations.Migrate()

	// Init database
	api.StartApi()
}
