package database

import (
	"ModaLast/src/models"
	"ModaLast/src/seeders"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDb() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env files")
	}

	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbUrl, dbPort, dbName)
	db, err := gorm.Open(dbDriver, dbString)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to %s", db)
	}

	// Drop tables
	db.DropTable(&models.User{})

	// Migrate tables
	db.AutoMigrate(&models.User{})

	seeders.LoadSeeders(db)

	return db, err
}
