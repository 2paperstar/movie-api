package service

import (
	"math/rand"

	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
	"github.com/2paperstar/movie-api/utils"
)

func convertMovie(movie database.Movie, detail bool) model.Movie {
	randomReviewCounts := rand.Intn(1000)
	data := model.Movie{
		ID:           movie.Imdb,
		Title:        movie.Title,
		Genre:        movie.Genre,
		PosterURL:    utils.IdToPosterUrl(movie.Identifier),
		ReviewCounts: randomReviewCounts,
		ReviewScore:  movie.Rating,
	}
	if len(movie.Director) > 0 {
		data.Director = &movie.Director[0]
	}
	if detail {
		data.Story = movie.Story
	}
	return data
}

func GetMovies() []model.Movie {
	movies := database.GetMoviesData()

	var movieList = make([]model.Movie, len(*movies))

	for index, movie := range *movies {
		movieList[index] = convertMovie(movie, false)
	}

	return movieList
}

func GetMovieDetail(id string) *model.Movie {
	movie := database.GetMovieByID(id)
	if movie == nil {
		return nil
	}
	data := convertMovie(*movie, true)
	return &data
}
