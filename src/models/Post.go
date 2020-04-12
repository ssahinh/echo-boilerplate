package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title       string `validate:"required"`
	Description string `validate:"required"`
	User        User
	UserId      uint `sql:"type:int REFERENCES users(id)"`
}
