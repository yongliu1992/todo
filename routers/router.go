package routers

import (
	"github.com/gin-gonic/gin"
	api "github.com/yongliu1992/todo/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiV1 := r.Group("/api/v1/")
	apiV1.GET("todo/:uid", api.FindsTodo)
	apiV1.GET("todoOne/:id", api.FindOneTodo)
	apiV1.PUT("todo/:id", api.UpdateTodo)
	apiV1.DELETE("todo/delete/:uid", api.DeleteTodo)
	apiV1.POST("todo/:uid", api.AddTodo)
	return r
}
