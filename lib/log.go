package lib

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	con "github.com/yongliu1992/todo/config"
	"path"
	//"runtime"
	"sync"
	"time"
)

//Log 日志对象
var Log *logrus.Logger

var once sync.Once
var lock *sync.Mutex = &sync.Mutex{}
var sysFields map[string]interface{}

// GetLogInstance 单例生成日志实例
/*
 * DebugLevel<InfoLevel<WarnLevel<ErrorLevel
 * 低级别显示比自己高级别的日志
 * 返回实例化对象
 */
func GetLogInstance() *logrus.Logger {

	var durationMaxAge = 24 * time.Hour
	var durationRotationTime = 24 * time.Hour
	baseLogPath := path.Join(con.LogPath, con.LogFileName)
	writer, _ := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(durationMaxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(durationRotationTime), // 日志切割时间间隔
	)
	once.Do(func() {
		Log = logrus.New()
		// 为不同级别设置不同的输出目的
		lfHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
		}, &logrus.JSONFormatter{})
		Log.AddHook(lfHook)
	})
	return Log
}
