package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string   `json:"title"`
	Image    string   `json:"image"`
	Synopsis string   `json:"synopsis"`
	Genres   []*Genre `json:"genres" gorm:"many2many:movie_genres;"`
}
