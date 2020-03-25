package config

import "os"

var (
	//MongoDatabase  DB configs MongoCollection
	MongoDatabase = "hello"
	//MongoCollection db collection
	MongoCollection = "todo"
	//LogPath 日志存放路径
	LogPath = os.Getenv("LOG_PATH")
	//LogFileName 日志文件名称
	LogFileName = os.Getenv("LOG_NAME")
	//JwtSecret jwt密钥
	JwtSecret = os.Getenv("JwtSecret")
)
