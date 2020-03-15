package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var DB *Database

//初始化
func Init() {
	DB = &Database{
		Mongo: SetConnect(),
	}

}

// 连接设置
func SetConnect() *mongo.Client {
	pwd := os.Getenv("MONGO_PASSWORD")
	user := os.Getenv("MONGO_USERNAME")
	host := os.Getenv("MONGO_HOST")
	uri := "mongodb+srv://%s:%s@%s" //此处根据实际替换
	uri = fmt.Sprintf(uri,user,pwd,host)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println("err",err)
		panic(err)
	}
	fmt.Println("err",client.Ping(context.Background(),nil))
	return client
}
