package seeders

import "github.com/jinzhu/gorm"

func LoadSeeders(db *gorm.DB) {
	UserSeeder(db)
	PostSeeder(db)
}
