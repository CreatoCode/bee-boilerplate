package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	engine *gin.Engine
}

type RouterGroup struct {
	routerGroup *gin.RouterGroup
}

var g_route *Engine

func init() {
	eng := gin.Default()
	g_route = &Engine{eng}
	// g_route.Use(gin.Recovery())
	// g_route.Use(gin.Logger())
	// router := gin.Default()
	// group := router.Group("router")
	// group.GET("aaa")
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })
}

func Default() *Engine {
	// r := gin.Default()
	// r.Use(gin.Recovery())
	// r.Use(gin.Logger())
	return g_route
}

func (e *Engine) Group(elativePath string) *RouterGroup {
	rg := e.engine.Group(elativePath)
	return &RouterGroup{routerGroup: rg}
}

func (e *Engine) Run(addr ...string) (err error) {
	return g_route.engine.Run(addr...)
}

func (g *RouterGroup) Get(relativePath string, handler HandlerFunc, model interface{}) {
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
	g.routerGroup.GET(relativePath, get)
}

func (g *RouterGroup) Post(relativePath string, handler HandlerFunc, model interface{}) {
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
	g.routerGroup.POST(relativePath, get)
}
