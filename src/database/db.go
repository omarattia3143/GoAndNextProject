package database

import (
	"GoAndNextProject/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {

	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("couldn't connect to database!")
	}

	log.Println("connected to database successfully!")

}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Link{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic("cannot migrate user model")
	}
}
