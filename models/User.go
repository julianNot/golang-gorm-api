package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	lastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique_index"`
	Tasks     []Task
}
