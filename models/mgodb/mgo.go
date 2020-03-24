package mgodb

import (
	"context"
	"fmt"
	"github.com/yongliu1992/todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mgo struct {
	database   string
	collection string
}

type Todo struct {
	Task       string `json:"task" bson:"task"`
	DueDate    string `json:"due_date" bson:"dueDate"`
	Labels     string `json:"labels" bson:"labels"`
	Comments   string `json:"comments" bson:"comments"`
	UID        int    `json:"uid" bson:"uid"`
	CreateTime string `json:"create_date" bson:"createDate"`
	UpdateTime string `json:"updateDate" bson:"updateDate"`
	ID         string `json:"id" bson:"_id,omitempty"`
	Status     int    `json:"status" bson:"status"` //1代表已完成
}

func init() {
	//如果单独测试里面的方法 会造成 models.DB为null 引发panic 所以在这里进行nil判断
	if models.DB == nil {
		models.Init()
	}
}

func NewMgo(database, collection string) *mgo {
	return &mgo{
		database,
		collection,
	}
}

// 查询单个
func (m *mgo) FindOne(id primitive.ObjectID) (res Todo, err error) {
	client := models.DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&res)
	return
}

//插入单个
func (m *mgo) InsertOne(value interface{}) (insertResult *mongo.InsertOneResult, err error) {

	client := models.DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	insertResult, err = collection.InsertOne(context.TODO(), value)
	if err != nil {
		fmt.Println(err)
		return
	}
	return insertResult, nil
}

//_id 删一个的话
func (m *mgo) Delete(key string, value interface{}) int64 {
	client := models.DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{primitive.E{Key: key, Value: value}}
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println(err)
		log.Print(err)
		return 0
	}
	return count.DeletedCount
}

//删除多个
func (m *mgo) DeleteMany(key string, value interface{}) int64 {
	client := models.DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{primitive.E{Key: key, Value: value}}

	count, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount
}

//_id 更新
func (m *mgo) Update(id interface{}, t Todo) (err error) {
	client := models.DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	set := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "task", Value: t.Task},
		primitive.E{Key: "dueDate", Value: t.DueDate},
		primitive.E{Key: "labels", Value: t.Labels},
		primitive.E{Key: "comments", Value: t.Comments},
		primitive.E{Key: "uid", Value: t.UID},
		primitive.E{Key: "updateDate", Value: t.UpdateTime},
		primitive.E{Key: "status", Value: t.Status},
	}}}
	err = collection.FindOneAndUpdate(context.TODO(), filter, set).Decode(&t)
	return
}

//查找多个
func (m *mgo) FindMany(key string, value interface{}, sort int) (data []Todo, err error) {

	client := models.DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{primitive.E{Key: key, Value: value}}
	sorts := bson.D{primitive.E{Key: "_id", Value: sort}}
	var cur *mongo.Cursor
	if key != "" && value != nil {
		cur, err = collection.Find(context.TODO(), filter, &options.FindOptions{Sort: sorts})
	} else {
		cur, err = collection.Find(context.TODO(), bson.D{}, &options.FindOptions{Sort: sorts})
	}
	if cur != nil {
		for cur.Next(context.Background()) {
			var t Todo
			cur.Decode(&t)
			data = append(data, t)
		}
	}

	return data, err
}
