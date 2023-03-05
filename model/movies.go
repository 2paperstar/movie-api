package model

import "time"

type Movie struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Director     *string  `json:"director"`
	Genre        []string `json:"genre"`
	PosterURL    string   `json:"poster_url"`
	ReviewCounts int      `json:"review_counts"`
	ReviewScore  float32  `json:"review_score"`
	Story        string   `json:"story,omitempty"`
}

type Comment struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movie_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
