// 注册路由
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/controller"
)

// 路由组
func RouterGroup(r *gin.Engine) {
	todoGroup := r.Group("todo")
	{
		todoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"router": "todo/index"})
		})
		todoGroup.POST("add", controller.TodoAdd)
		todoGroup.PUT("edit", controller.TodoEdit)
	}

}
