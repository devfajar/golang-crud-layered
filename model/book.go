package model

import "time"

type Book struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Author         string    `json:"author"`
	Published_Date *time.Time `json:"published_date"`
}
