package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/config"
	"github.com/yongliu1992/todo/models/mgodb"
	"github.com/yongliu1992/todo/pkg/e"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var dbName = config.MongoDatabase
var coll = config.MongoCollection

func Delete(c *gin.Context) {
	g := Gin{C: c}
	uid := g.PostInt("uid")
	mgo := mgodb.NewMgo(dbName, coll)
	if uid < 1 {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	keyName := c.PostForm("key")
	value := c.PostForm("value")
	var rows int64
	if keyName == "uid"{
		rows = mgo.DeleteMany(keyName, uid)
	}else {
		rows = mgo.DeleteMany(keyName, value)
	}


	res := make(map[string]interface{}, 1)
	res["data"] = rows
	g.Response(e.SUCCESS, res)

}

func Update(c *gin.Context) {

}

func Add(c *gin.Context) {
	g := Gin{C: c}
	uid := g.PostInt("uid")
	mgo := mgodb.NewMgo(dbName, coll)
	if uid < 1 {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	task := c.PostForm("task")
	DueDate := c.PostForm("endDate")
	Labels := c.PostForm("label")
	Comment := c.PostForm("comm")
	data := mgodb.Todo{
		Task:       task,
		DueDate:    DueDate,
		Labels:     Labels,
		Comments:   Comment,
		Uid:        uid,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: "",
	}
	dataS, err := mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		g.Response(e.ERROR, map[string]interface{}{"err": err})
	} else {
		res := make(map[string]interface{}, 1)
		res["data"] = dataS
		g.Response(e.SUCCESS, res)
	}
}

/**
 * @api {Get} /todo/index 获取todo
 * @apiName api.GetLists
 * @apiGroup todo_lists
 * @apiParam {Int} uid 用户uid
 * @apiParam {Int} [sort] 排序
 * @apiSuccess {int} code 错误码，0-成功.
 * @apiSuccess {String} error 错误信息.
 * @apiSuccess {list} data 数据集合.
 */
func Finds(c *gin.Context) {
	g := Gin{C: c}
	uid := g.GetInt("uid")
	sort := g.GetInt("sort")
	if uid < 1 {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	mgo := mgodb.NewMgo(dbName, coll)
	dataS, err := mgo.FindMany("uid", uid, sort)
	res := make(map[string]interface{}, 2)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			res["data"] = map[string]interface{}{}
			res["count"] = 0
			g.Response(e.SUCCESS, res)
			return
		}
		g.Response(e.ERROR, map[string]interface{}{"err": err})
	} else {
		res["data"] = dataS
		res["count"] = len(dataS)
		g.Response(e.SUCCESS, res)
	}
}
