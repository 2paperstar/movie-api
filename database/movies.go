package database

import (
	_ "embed"
	"encoding/json"
)

type Movie struct {
	Imdb     string   `json:"imdb"`
	Status   string   `json:"status"`
	Title    string   `json:"title"`
	Year     string   `json:"year"`
	Runtime  string   `json:"runtime"`
	Story    string   `json:"story"`
	Director []string `json:"director"`
	Writers  []string `json:"writers"`
	Stars    []string `json:"stars"`
	Genre    []string `json:"genre"`
	Rating   float32  `json:"rating,string"`
	Related  []string `json:"related"`
	AoMeta   struct {
		Subject []string `json:"subject"`
	} `json:"aoMeta"`
	AoFiles []struct {
		URL    string `json:"url"`
		Format string `json:"format"`
	} `json:"aoFiles"`
	Identifier           string `json:"identifier"`
	ArchiveDownloadCount int    `json:"archiveDownloadCount"`
	Wiki                 string `json:"wiki"`
	StoryFrom            string `json:"storyFrom"`
}

//go:embed movies.json
var moviesRawData []byte

func GetMovieData() *[]Movie {
	var movies []Movie
	json.Unmarshal(moviesRawData, &movies)
	return &movies
}
