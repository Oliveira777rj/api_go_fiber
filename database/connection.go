package database

import (
	"api_fiber/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	MYSQLDATABASE := os.Getenv("MYSQLDATABASE")
	MYSQLHOST := os.Getenv("MYSQLHOST")
	MYSQLPASSWORD := os.Getenv("MYSQLPASSWORD")
	MYSQLPORT := os.Getenv("MYSQLPORT")
	MYSQLUSER := os.Getenv("MYSQLUSER")

	dsn := MYSQLUSER +":"+ MYSQLPASSWORD +"@tcp"+ "(" + MYSQLHOST + ":" + MYSQLPORT +")/" + MYSQLDATABASE + "?" + "parseTime=true&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
