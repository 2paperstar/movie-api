package controller

import (
	"github.com/2paperstar/movie-api/service"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"movies": service.GetMovies(),
	})
}

func GetMovieDetail(c *gin.Context) {
	movie := service.GetMovieDetail(c.Param("id"))
	if movie == nil {
		c.JSON(404, gin.H{"message": "movie not found"})
		return
	}
	c.JSON(200, movie)
}
