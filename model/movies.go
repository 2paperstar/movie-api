package model

import "time"

type Movie struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Director     *string  `json:"director"`
	Genre        []string `json:"genre"`
	PosterURL    string   `json:"posterUrl"`
	ReviewCounts int      `json:"reviewCounts"`
	ReviewScore  float32  `json:"reviewScore"`
	Story        string   `json:"story,omitempty"`
}

type Comment struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movieId"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
