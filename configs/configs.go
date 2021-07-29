package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"test/project/models"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
