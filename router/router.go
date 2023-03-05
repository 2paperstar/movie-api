package router

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) *gin.Engine {
	movies(r.Group("/movies"))
	return r
}
