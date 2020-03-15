package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yongliu1992/todo/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiV1 := r.Group("/api/v1/")
	apiV1.GET("todo/index", v1.Finds)
	apiV1.POST("todo/update", v1.Update)
	apiV1.POST("todo/delete", v1.Delete)
	apiV1.POST("todo/save", v1.Add)
	return r
}
