package main

import (
	_ "github.com/2paperstar/movie-api/docs"
	"github.com/2paperstar/movie-api/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Movie API
// @description This is a sample server for a movie API.
// @version 1.0.1

// @host test.paperst.ar
func setup() *gin.Engine {
	r := router.SetupRouter()
	r.Use(cors.Default())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
