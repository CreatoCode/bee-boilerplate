package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine
}

type RouteGroup struct {
	router *gin.RouterGroup
}

var s_route *Engine

func init() {
	eng := gin.Default()
	s_route = &Engine{eng}
	// s_route.Use(gin.Recovery())
	// s_route.Use(gin.Logger())
	// router := gin.Default()
	// group := router.Group("router")
	// group.GET("aaa")
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })
}

func Default() *IEngine {
	// r := gin.Default()
	// r.Use(gin.Recovery())
	// r.Use(gin.Logger())
	return s_route
}

func (e *Engine) Group(elativePath string) *IRouterGroup {
	return e.Group(elativePath)
}

func (g *RouteGroup) GET(relativePath string, handler HandlerFunc, model any) {
	get := func(gc *gin.Context) {
		if err := gc.ShouldBindJSON(&model); err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c := &Context{model}
		if data, err := handler(c); err != nil {
			gc.JSON(http.StatusOK, ResponseData{Msg: "ok", Code: "0", Data: data})
		}
	}
	g.router.GET(relativePath, get)
}
