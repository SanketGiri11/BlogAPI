package database

import (
	"blogapi/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//Connect the DB with sqlite DB and the DB file name is sqlite.db
	database, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to Database")
	} else {
		fmt.Println("DB connected::", database)
	}
	DB = database

	//Migrating all the data to the specified table
	DB.AutoMigrate(&models.User{}, &models.Posts{}, &models.Comments{})

}
