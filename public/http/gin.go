package http

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	return r
}
