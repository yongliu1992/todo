package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yongliu1992/todo/config"
	"github.com/yongliu1992/todo/models/mgodb"
	"github.com/yongliu1992/todo/pkg/e"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

var dbName = config.MongoDatabase
var coll = config.MongoCollection

func DeleteTodo(c *gin.Context) {
	g := Gin{C: c}
	uid, _ := strconv.Atoi(c.Param("uid"))
	mgo := mgodb.NewMgo(dbName, coll)
	keyName := c.Query("key")
	value := c.Query("value")
	if uid < 1 || keyName == "" || value == "" {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	var rows int64
	if keyName == "uid" {
		rows = mgo.DeleteMany(keyName, uid)
	} else {
		rows = mgo.DeleteMany(keyName, value)
	}
	g.Response(e.SUCCESS, rows)
}

func UpdateTodo(c *gin.Context) {
	g := Gin{C: c}
	uid := g.PostInt("uid")
	id := c.Param("id")
	mgo := mgodb.NewMgo(dbName, coll)
	if uid < 1 {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	task := c.PostForm("task")
	DueDate := c.PostForm("endDate")
	Labels := c.PostForm("label")
	Comment := c.PostForm("comm")
	status := g.PostInt("status")
	data := mgodb.Todo{
		Task:       task,
		DueDate:    DueDate,
		Labels:     Labels,
		Comments:   Comment,
		Uid:        uid,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		Status:     status,
	}
	tid, _ := primitive.ObjectIDFromHex(id)
	err := mgo.Update(tid, data)
	if err != nil {
		fmt.Println("error", err)
		g.Response(e.ERROR, map[string]interface{}{"err": err})
	} else {
		g.Response(e.SUCCESS, map[string]interface{}{})
	}
}

func AddTodo(c *gin.Context) {
	g := Gin{C: c}
	uid, _ := strconv.Atoi(c.Param("uid"))
	mgo := mgodb.NewMgo(dbName, coll)
	if uid < 1 {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	task := c.PostForm("task")
	DueDate := c.PostForm("endDate")
	Labels := c.PostForm("label")
	Comment := c.PostForm("comm")
	status := g.PostInt("status")
	data := mgodb.Todo{
		Task:       task,
		DueDate:    DueDate,
		Labels:     Labels,
		Comments:   Comment,
		Uid:        uid,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: "",
		Status:     status,
	}
	dataS, err := mgo.InsertOne(data)
	if err != nil {
		g.Response(e.ERROR, "")
	} else {
		g.Response(e.SUCCESS, dataS)
	}
}

/**
 * @api {Get} /todo/index 获取todo
 * @apiName api.GetLists
 * @apiGroup todoP
 * @apiParam {Int} uid 用户uid
 * @apiParam {Int} [sort] 排序
 * @apiSuccess {int} code 错误码，0-成功.
 * @apiSuccess {String} error 错误信息.
 * @apiSuccess {list} data 数据集合.
 */
func FindsTodo(c *gin.Context) {
	g := Gin{C: c}
	uid, _ := strconv.Atoi(c.Param("uid"))
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
			res["data"] = []int{}
			res["count"] = 0
			g.Response(e.SUCCESS, res)
			return
		}
		g.Response(e.ERROR, map[string]interface{}{"err": err})
	} else {
		res["data"] = dataS
		res["count"] = len(dataS)
		if res["count"] == 0 {
			res["data"] = []int{}
		}
		g.Response(e.SUCCESS, res)
	}
}

/**
 * @api {Get} /todoOne/:id 获取todo
 * @apiName api.FindOneTodo
 * @apiGroup todoP
 * @apiParam {Int} uid 用户uid
 * @apiParam {Int} [sort] 排序
 * @apiSuccess {int} code 错误码，0-成功.
 * @apiSuccess {String} error 错误信息.
 * @apiSuccess {list} data 数据集合.
 */
func FindOneTodo(c *gin.Context) {
	g := Gin{C: c}
	id := c.Param("id")
	if id == "" {
		g.Response(e.ERROR_PARAM_ERROR, map[string]interface{}{})
		return
	}
	mgo := mgodb.NewMgo(dbName, coll)
	tid, _ := primitive.ObjectIDFromHex(id)
	dataS, err := mgo.FindOne(tid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			g.Response(e.SUCCESS, map[string]interface{}{})
			return
		}
		g.Response(e.ERROR, map[string]interface{}{})
	} else {

		g.Response(e.SUCCESS, dataS)
	}
}
