package config

import "os"

var (
	// DB configs
	MongoDatabase   = "hello"
	MongoCollection = "todo"
	LogPath         = os.Getenv("LOG_PATH")
	LogFileName     = os.Getenv("LOG_NAME")
	JwtSecret       = os.Getenv("JwtSecret")
)
