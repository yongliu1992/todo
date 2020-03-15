package config

import "os"

var (
    // DB configs
    MongoHost = os.Getenv("MONGO_HOST")
    MongoPort = os.Getenv("MONGO_PORT")
    MongoUsername = os.Getenv("MONGO_USERNAME")
    MongoPassword = os.Getenv("MONGO_PASSWORD")
    //MongoDatabase = os.Getenv("MONGO_DATABASE")
    //MongoCollection = os.Getenv("MONGO_COLLECTION")
    MongoDatabase = "hello"
    MongoCollection = "todo"
    TodoServiceIP = os.Getenv("TODO_SERVICE_IP")
    TodoServicePort = os.Getenv("TODO_SERVICE_PORT")
    LogPath = os.Getenv("LOG_PATH")
    LogFileName = os.Getenv("LOG_NAME")
    JwtSecret = os.Getenv("JwtSecret")

)
