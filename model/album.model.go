package model

import (
	"gorm.io/gorm"
)

// album represents data about a record album.
type Album struct {
	gorm.Model         // https://gorm.io/docs/models.html
	Title      string  `json:"title"`
	Artist     string  `json:"artist"`
	Price      float64 `json:"price"`
}
