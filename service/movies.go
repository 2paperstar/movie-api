package service

import (
	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
)

func GetMovies() *[]model.Movie {
	const CDN_BASE = "https://boomerangdb.nyc3.cdn.digitaloceanspaces.com/"
	movies := database.GetMovieData()

	var movieList = make([]model.Movie, len(*movies))

	for index, movie := range *movies {
		movieList[index] = model.Movie{
			ID:           movie.Imdb,
			Title:        movie.Title,
			Director:     movie.Director[0],
			Genre:        movie.Genre,
			PosterURL:    CDN_BASE + "posters/webp/" + movie.Identifier + ".webp",
			ReviewCounts: 0,
			ReviewScore:  movie.Rating,
		}
	}

	return &movieList
}
