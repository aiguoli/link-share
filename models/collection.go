package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	ShareID uint `json:"share_id"`
}
