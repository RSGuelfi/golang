package migrations

import (
	"github.com/RSGuelfi/golang/bankingApp/database"
	"github.com/RSGuelfi/golang/bankingApp/helpers"
	"github.com/RSGuelfi/golang/bankingApp/interfaces"
	_ "github.com/lib/pq"
)

// Refactor createAccounts to use database package
func createAccounts() {
	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}
	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

// Refactor Migrate
func Migrate() {
	database.DB.AutoMigrate(&interfaces.User{}, &interfaces.Account{}, &interfaces.Transaction{})

	createAccounts()
}
