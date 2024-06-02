package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string
	Body      string
}
