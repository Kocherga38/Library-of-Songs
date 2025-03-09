package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Group string `json:"group"`
	Song  string `json:"song" gorm:"unique"`
}
