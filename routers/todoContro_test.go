package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yongliu1992/todo/models/mgodb"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

var router = InitRouter()

//
type TodoListResponse struct {
	Code  int `json:"code"`
	Data  TodoListResponseData
	Error string `json:"error"`
}
type TodoDetailResponse struct {
	Code  int `json:"code"`
	Data  mgodb.Todo
	Error string `json:"error"`
}

type TodoListResponseData struct {
	Data  []mgodb.Todo `json:"data"`
	Count int          `json:"count"`
}
type InsertDetailTodoResponseData struct {
	InsertedID string
}
type InsertTodoResponseData struct {
	Data  InsertDetailTodoResponseData `json:"data"`
	Code  int                          `json:"code"`
	Error string                       `json:"error"`
}

type TodoIntResponse struct {
	Code  int    `json:"code"`
	Data  int    `json:"data"`
	Error string `json:"error"`
}

func TestAddTodo(t *testing.T) {
	var param = url.Values{}
	rand.Seed(time.Now().Unix())
	param.Add("task", "task"+strconv.Itoa(rand.Intn(1000)))
	param.Add("label", "label"+strconv.Itoa(rand.Intn(1000)))
	param.Add("comm", "comm")
	param.Add("status", "1")
	param.Add("endDate", time.Now().Format("2006-01-02 15:04:05"))
	dataByte := PostForm("/api/v1/todo/1", param, router)
	var res InsertTodoResponseData
	err := json.Unmarshal(dataByte, &res)
	assert.Nil(t, err, err)
}

func TestFindsTodo(t *testing.T) {
	url := "/api/v1/todo/1?sort=1"
	data := Get(url, router)
	var res TodoListResponse
	fmt.Println(string(data))
	err := json.Unmarshal(data, &res)
	assert.Nil(t, err, err)

}

//
func TestDeleteTodo(t *testing.T) {
	url := "/api/v1/todo/1?sort=1"
	data := Get(url, router)
	var res TodoListResponse
	err := json.Unmarshal(data, &res)
	assert.Nil(t, err, err)
	lastIndex := len(res.Data.Data)
	lastTask := res.Data.Data[lastIndex-1].Task
	urlD := "/api/v1/todo/delete/1?key=task&value=" + lastTask
	dataByte := Delete(urlD, router)
	var res2 TodoIntResponse
	err = json.Unmarshal(dataByte, &res2)
	assert.Nil(t, err, err)
	assert.Equal(t, 1, res2.Data)

	url = "/api/v1/todo/1?sort=1"
	data = Get(url, router)
	err = json.Unmarshal(data, &res)
	assert.Nil(t, err, err)
	lastIndex = len(res.Data.Data)
	if lastIndex > 0 {
		assert.NotEqual(t, lastTask, res.Data.Data[lastIndex-1].Task)
	}
}

func TestUpdateRule(t *testing.T) {
	apiUrl := "/api/v1/todo/1?sort=1"
	data := Get(apiUrl, router)
	var res TodoListResponse
	err := json.Unmarshal(data, &res)
	assert.Nil(t, err, err)
	lastIndex := len(res.Data.Data)
	lastData := res.Data.Data[lastIndex-1]
	rand.Seed(time.Now().Unix())
	rand.Intn(1000)

	var param = url.Values{}
	newTask := "task_" + strconv.Itoa(rand.Intn(1000))
	param.Add("task", newTask)
	param.Add("label", "label"+strconv.Itoa(rand.Intn(1000)))
	param.Add("comm", "comm")
	param.Add("status", "1")
	param.Add("endDate", time.Now().Format("2006-01-02 15:04:05"))
	param.Add("uid", strconv.Itoa(lastData.Uid))
	PutForm("/api/v1/todo/"+lastData.Id, param, router)
	dataByte := Get("/api/v1/todoOne/"+lastData.Id, router)
	fmt.Println(string(dataByte))
	var res2 TodoDetailResponse
	err = json.Unmarshal(dataByte, &res2)
	assert.Nil(t, err, err)
	assert.Equal(t, newTask, res2.Data.Task)
}

func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func Delete(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("DELETE", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应的handler接口
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param url.Values, router *gin.Engine) []byte {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PutForm(uri string, param url.Values, router *gin.Engine) []byte {
	// 构造post请求
	req := httptest.NewRequest("PUT", uri, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	router.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
