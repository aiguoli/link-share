package models

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	gorm.Model
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Password       string    `json:"password"`
	ExpirationDate time.Time `json:"expiration_date"`
	UserID         uint      `json:"user_id"`
	Views          uint      `json:"views"`
}
