package controller

import (
	"strconv"

	"github.com/2paperstar/movie-api/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get all movies
// @Description Get all movies
// @Tags movies
// @Produce  json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} model.Empty{movies=[]model.Movie}
// @Router /movies [get]
func GetMovies(c *gin.Context) {
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

	c.JSON(200, gin.H{
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
func GetMovieDetail(c *gin.Context) {
	movie := service.GetMovieDetail(c.Param("id"))
	if movie == nil {
		c.JSON(404, gin.H{"message": "movie not found"})
		return
	}
	c.JSON(200, movie)
}
