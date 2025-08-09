package models

import "time"

type Blog struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Slug        string    `json:"slug" gorm:"not null;unique"`
	Content     string    `json:"content" gorm:"not null"`
	Image       string    `json:"image" gorm:"not null"`
	ReadingTime int       `json:"readingTime" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;"`
}
