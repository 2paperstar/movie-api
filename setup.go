package main

import (
	"github.com/2paperstar/movie-api/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Movie API
// @description This is a sample server for a movie API.
// @version 1.0.0

// @host https://test.com:8080
func setup() *gin.Engine {
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}