package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/controller"
	"go-web/router"
)

var Router *gin.Engine

func RegisRouter() {
	Router = gin.Default()
	Router.Use(controller.StatCost())

	Router.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.RouterGroup(Router)
}
