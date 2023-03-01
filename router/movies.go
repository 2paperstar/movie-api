package router

import (
	"github.com/2paperstar/movie-api/controller"
	"github.com/gin-gonic/gin"
)

func movies(r *gin.RouterGroup) {
	r.GET("/", controller.GetMovies)
}
