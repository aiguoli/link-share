package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string       `gorm:"uniqueIndex;not null" json:"username"`
	Email       string       `gorm:"uniqueIndex;not null" json:"email"`
	Password    string       `json:"password"`
	Links       []Link       `json:"links"`
	Collections []Collection `gorm:"many2many:collections" json:"collections"`
}
