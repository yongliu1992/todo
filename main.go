package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yongliu1992/todo/lib"
	"github.com/yongliu1992/todo/pkg/util"
	"github.com/yongliu1992/todo/routers"
)

var _version_ = ""
var _branch_ = ""
var _commitId_ = ""
var _userName_ = ""
var _buildTime_ = ""

var logger = lib.GetLogInstance()

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "-v")
	flag.Parse()

	if version {
		fmt.Printf("Version: %s, Branch: %s, Build: %s, User: %s, Build time: %s\n", _version_, _branch_, _commitId_, _userName_, _buildTime_)
	} else {
		util.Setup()
		logger.SetLevel(logrus.ErrorLevel)
		r := routers.InitRouter()
		r.Run(":8080")
	}
}
