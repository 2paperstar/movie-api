package router

import "github.com/gin-gonic/gin"

func movies(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": []string{"1", "2"},
		})
	})
}
