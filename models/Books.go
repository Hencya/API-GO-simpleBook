package Models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	ID        int    `json:"ID"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer"`
}