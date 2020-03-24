package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yongliu1992/todo/lib"
	"github.com/yongliu1992/todo/pkg/util"
	"github.com/yongliu1992/todo/routers"
)

var _version = ""
var _branch = ""
var _commitID = ""
var _userName = ""
var _buildTime = ""

var logger = lib.GetLogInstance()

/*TODO: Need front-end UI to make a beautiful product */
func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "-v")
	flag.Parse()

	if version {
		fmt.Printf("Version: %s, Branch: %s, Build: %s, User: %s, Build time: %s\n", _version, _branch, _commitID, _userName, _buildTime)
	} else {
		util.Setup()
		logger.SetLevel(logrus.ErrorLevel)
		r := routers.InitRouter()
		r.Run("127.0.0.1:8080")
	}
}
