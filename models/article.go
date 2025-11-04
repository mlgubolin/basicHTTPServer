package models

import "gorm.io/gorm"

// Article model
type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug" gorm:"uniqueIndex"`
	Author  string `json:"author"`
}
