package models

import "time"

type Task struct {
	Id          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Due         time.Time
	CreatedAt   time.Time
	Open        bool `gorm:"default:true"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
}
