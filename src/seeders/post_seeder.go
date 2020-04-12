package seeders

import (
	"ModaLast/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

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

func PostSeeder(db *gorm.DB) {
	for i, _ := range posts {
		posts[i].User = users[i]
		err := db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		err = db.Debug().Model(&models.User{}).Where("id = ?", users[i].ID).Take(&posts[i].User).Error
		fmt.Println(&posts[i])
		if err != nil {
			log.Fatalf("cannot seed posts, %v", err)
		}
	}
}
