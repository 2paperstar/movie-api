package controller

import (
	"strconv"
	"time"

	"github.com/2paperstar/movie-api/model"
	"github.com/2paperstar/movie-api/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all movies
// @Description Get all movies
// @Tags movies
// @Produce  json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} model.Empty{movies=[]model.Movie}
// @Router /movies [get]
func GetMovies(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	if limit == 0 {
		limit = 10
	}

	movies := service.GetMovies()

	if offset > len(movies) {
		offset = len(movies)
	}

	if offset+limit > len(movies) {
		limit = len(movies) - offset
	}

	return c.JSON(fiber.Map{
		"movies": movies[offset : offset+limit],
	})
}

// @Summary Get movie detail
// @Description Get movie detail (including story)
// @Tags movies
// @Produce  json
// @Param id path string true "Movie ID"
// @Success 200 {object} model.Movie
// @Failure 404 {object} model.Error
// @Router /movies/{id} [get]
func GetMovieDetail(c *fiber.Ctx) error {
	movie := service.GetMovieDetail(c.Params("id"))
	if movie == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "movie not found",
		})
	}
	return c.JSON(movie)
}

// @Summary Get movie comments
// @Description Get comments of movies
// @Tags movies
// @Produce json
// @Param id path string true "Movie ID"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} model.Empty{comments=[]model.Comment}
// @Failure 404 {object} model.Error
// @Router /movies/{id}/comments [get]
func GetMovieComments(c *fiber.Ctx) error {
	movie := service.GetMovieDetail(c.Params("id"))
	limit := c.QueryInt("limit", 10)
	c.QueryInt("offset", 0)
	if movie == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "movie not found",
		})
	}

	comments := make([]model.Comment, limit)

	for i := 0; i < limit; i++ {
		comments[i] = model.Comment{
			ID:        strconv.Itoa(i),
			MovieID:   movie.ID,
			Content:   "This is a comment",
			Author:    "eeee",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
	}

	return c.JSON(fiber.Map{
		"comments": comments,
	})
}
