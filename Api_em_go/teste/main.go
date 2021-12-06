package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID        int
	Uname     string
	Languages []Language `gorm:"many2many:user_languages";"ForeignKey:UserId"`
}

type Language struct {
	ID   int
	Name string
}

type UserLanguages struct {
	UserId     int
	LanguageId int
}

var db *gorm.DB

func initDatabase() {

	var err error
	db, err = gorm.Open("sqlite3", "owners.db")
	if err != nil {
		panic("Error connecting to database.")
	}
	db.AutoMigrate(&User{}, &Language{}, &UserLanguages{})
	fmt.Println("Database Connected...")
	fmt.Println("Database Migrated")

}

func main() {
	//db.LogMode(true)

	defer db.Close()

	//db.DropTableIfExists(&Contact{},&Customer{})
	db.SingularTable(true)
	// generate a database table

	//db.Model(UserLanguages{}).AddForeignKey("user_id","user(id)","CASCADE","CASCADE")
	//db.Model(UserLanguages{}).AddForeignKey("language_id","language(id)","CASCADE","CASCADE")

	langs := []Language{{Name: "English"}, {Name: "Janpanies"}, {Name: "Franch"}}

	user1 := User{Uname: "ChenY", Languages: langs}
	user2 := User{Uname: "LUS", Languages: langs}

	// add cascade
	db.Save(&user1)
	db.Save(&user2)

	fmt.Println(&user1)
	fmt.Println(&user2)

	// cascade inquiry
	user := &User{}
	db.Where("id = ?", 2).Preload("Languages").Find(&user)
	fmt.Println(user)

	// delete cascade is not implemented
	// can only delete the master table data. Seek the guidance of the great God.
	db.Debug().Delete(&user)

	// cascading updates unrealized
	// only update the main table. Seek the guidance of the great God. . .
	user3 := User{Uname: "UpUser", Languages: []Language{{Name: "Shenss1"}, {Name: "Shen2"}}}

	db.Model(&User{}).Where("id = ?", 1).Update(user3)
	initDatabase()
}
