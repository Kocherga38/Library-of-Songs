package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Group string `json:"group"`
	Title string `json:"title"`
}
