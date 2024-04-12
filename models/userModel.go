package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Phone      string `json:"phone"`
}

type Word struct {
	gorm.Model
	Term string `json:"term"`
}

type Synonym struct {
	gorm.Model
	Word_id1 int16 `json:"word_id1"`
	Word_id2 int16 `json:"word_id2"`
}

type SearchResult struct {
	ID    int16  `json:"id"`
	Term  string `json:"term"`
	Level int16  `json:"level"`
}
