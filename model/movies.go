package model

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
