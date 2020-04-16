package test

import (
	"ModaLast/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BaseTest() *gorm.DB {
	db, err := connectDb()
	if err != nil {
		log.Fatal("db", err)
	}

	for i, _ := range users {
		users[i].Hash(users[i].Password)
		err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range posts {
		posts[i].User = users[i]
		err := db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		err = db.Debug().Model(&models.User{}).Where("id = ?", users[i].ID).Take(&posts[i].User).Error
		fmt.Println(&posts[i])
		if err != nil {
			log.Fatalf("cannot seed posts, %v", err)
		}
	}

	return db
}

func connectDb() (*gorm.DB, error) {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "localhost", "3306", "go_test")
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		fmt.Printf("Cannot connect to mysql database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("Connected to database")
	}

	// Drop tables
	db.DropTable(&models.User{}, &models.Post{})

	// Migrate tables
	db.AutoMigrate(&models.User{}, &models.Post{})

	return db, err
}

var posts = []models.Post{
	models.Post{
		Title:       "Post Seeder Title1",
		Description: "Post Description Example",
	},
	models.Post{
		Title:       "Post Seeder Title2",
		Description: "Post Description Example2",
	},
}

var users = []models.User{
	models.User{
		FullName: "Mike Hunt",
		Email:    "testuser@test.com",
		Password: "password",
	},
	models.User{
		FullName: "Hunter Mike",
		Email:    "test2user@test.com",
		Password: "password",
	},
}
