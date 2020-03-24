package mgodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestMgo_InsertOne(t *testing.T) {
	mgo := NewMgo("hello", "todoTest")

	data := Todo{
		Task:     "终生学习",
		DueDate:  "lastDay",
		Labels:   "日拱一卒",
		Comments: "书中自有颜如玉，书中自有黄金屋，与高尚快乐作伴",
		Uid:      1,
	}
	inResult, err := mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}
	fmt.Println("insert id", inResult.InsertedID)
	res, _ := mgo.FindOne(inResult.InsertedID.(primitive.ObjectID))
	fmt.Println("res", res.Comments, res.Task, res.Labels, res.DueDate, res.Uid)
}

func TestMgo_FindOne(t *testing.T) {
	mgo := NewMgo("hello", "todoTest")
	tid, _ := primitive.ObjectIDFromHex("5e6e143645b93c602616fe9e")
	res, err := mgo.FindOne(tid)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}
	fmt.Println("res", res.Comments, res.Task, res.Labels, res.DueDate)
}

func TestMgo_Delete(t *testing.T) {
	time.Sleep(4 * time.Second)
	mgo := NewMgo("hello", "todoTest")
	data := Todo{
		Task:     "终生学习",
		DueDate:  "lastDay",
		Labels:   "日拱一卒",
		Comments: "书中自有颜如玉，书中自有黄金屋，与高尚快乐作伴",
		Uid:      1,
	}
	inResult, err := mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}
	affectRows := mgo.Delete("_id", inResult.InsertedID)
	fmt.Println("del", affectRows)
	if affectRows < 1 {
		t.Error("删除行数", affectRows, " ")
		t.Fail()
	}
	_, err = mgo.FindOne(inResult.InsertedID.(primitive.ObjectID))
	if err != mongo.ErrNoDocuments {
		t.Error("删除失败")
		t.Fail()
	}
}

func TestMgo_Update(t *testing.T) {
	mgo := NewMgo("hello", "todoTest")
	data := Todo{
		Task:     "终生学习",
		DueDate:  "lastDay",
		Labels:   "日拱一卒",
		Comments: "书中自有颜如玉，书中自有黄金屋，与高尚快乐作伴",
		Uid:      1,
	}
	inResult, err := mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}

	data.Uid = 2
	data.Task = "学习至上"
	data.DueDate = "every day is first day"
	data.Labels = "滴水穿石"
	data.Comments = "将军府"

	err = mgo.Update(inResult.InsertedID, data)
	fmt.Println("update has err === ", err)
	if err != nil {
		t.Error("更新失败", err.Error(), " ")
		t.Fail()
	}
	mgo.Delete("_id", inResult.InsertedID)
}

func TestMgo_FindMany(t *testing.T) {
	mgo := NewMgo("hello", "todoTestMany")
	data := Todo{
		Task:     "终生学习",
		DueDate:  "lastDay",
		Labels:   "日拱一卒",
		Comments: "书中自有颜如玉，书中自有黄金屋，与高尚快乐作伴",
		Uid:      1,
	}
	_, err := mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}
	data = Todo{
		Task:     "终生学习",
		DueDate:  "lastDay",
		Labels:   "日拱一卒",
		Comments: "书中自有颜如玉，书中自有黄金屋，与高尚快乐作伴",
		Uid:      2,
	}
	_, err = mgo.InsertOne(data)
	if err != nil {
		fmt.Println("error", err)
		t.Log("err", err)
		t.Fail()
	}

	dataS, err := mgo.FindMany("dueDate", "lastDay", 1)
	if err != nil {
		t.Error("更新失败", err.Error(), " ")
		t.Fail()
	}
	if len(dataS) != 2 {
		t.Error("查找失败 数量不正确", len(dataS))
		t.Fail()
	}
	fmt.Println("TestMgo_FindMany deleted rows", mgo.DeleteMany("dueDate", "lastDay"))
}
