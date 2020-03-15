package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/pkg/e"
	"net/url"
	"strconv"
)

type Gin struct {
	C *gin.Context
	formCache url.Values
	engine    *gin.Engine
}

func (g *Gin) Response(errCode int, data interface{}) {
	g.C.JSON(200, gin.H{
		"code":  errCode,
		"error": e.GetMsg(errCode),
		"data":  data,
	})
	return
}

func (g *Gin) PostInt(pName string ) int  {
	p := g.C.PostForm(pName)
	pi,_ := strconv.Atoi(p)
	return pi
}

func (g *Gin) GetInt(pName string ) int  {
	p := g.C.Query(pName)
	pi,_ := strconv.Atoi(p)
	return pi
}