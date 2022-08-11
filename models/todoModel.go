package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Item      string
	Completed bool
	UserID    int64
}
