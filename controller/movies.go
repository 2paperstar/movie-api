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
