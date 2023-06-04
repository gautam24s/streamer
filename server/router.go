package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.NoRoute()
	router.NoMethod()

	return router
}