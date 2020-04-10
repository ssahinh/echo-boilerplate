package seeders

import (
	"ModaLast/src/models"
	"github.com/jinzhu/gorm"
	"log"
)

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

func RunUserSeeder(db *gorm.DB) {
	for i, _ := range users {
		users[i].Hash(users[i].Password)
		err := db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
