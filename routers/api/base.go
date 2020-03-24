package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/pkg/e"
	"net/url"
	"strconv"
)
// Gin 包装gin
type Gin struct {
	C         *gin.Context
	formCache url.Values
	engine    *gin.Engine
}

// Response 统一返回
func (g *Gin) Response(errCode int, data interface{}) {
	g.C.JSON(200, gin.H{
		"code":  errCode,
		"error": e.GetMsg(errCode),
		"data":  data,
	})
	return
}

//PostInt 转类型
func (g *Gin) PostInt(pName string) int {
	p := g.C.PostForm(pName)
	pi, _ := strconv.Atoi(p)
	return pi
}

//GetInt get query 参数转类型
func (g *Gin) GetInt(pName string) int {
	p := g.C.Query(pName)
	pi, _ := strconv.Atoi(p)
	return pi
}
