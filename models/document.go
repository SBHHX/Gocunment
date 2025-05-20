package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	Title       string `gorm:"not null" json:"title"`
	ContentPath string `gorm:"not null" json:"content_path"`
	IsPublic    bool   `gorm:"default:false" json:"is_public"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
}
