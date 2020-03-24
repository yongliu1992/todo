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
	MongoHost := os.Getenv("MONGO_HOST")
	MongoPort := os.Getenv("MONGO_PORT")
	MongoUsername := os.Getenv("MONGO_USERNAME")
	MongoPassword := os.Getenv("MONGO_PASSWORD")
	//uri := "mongodb://root:root@localhost:27017/"
	uri := "mongodb://%s:%s@%s:%s/mydb_test?gssapiServiceName=mongodb"
	uri = fmt.Sprintf(uri, MongoUsername, MongoPassword, MongoHost, MongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	fmt.Println("err", client.Ping(context.Background(), nil))
	return client
}
