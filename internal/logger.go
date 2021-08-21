package internal

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var Logger = logrus.New() // logrus 实例

var Log *logrus.Entry       // 默认日志
var LogAccess *logrus.Entry // 请求日志

func init() {
	Logger.SetReportCaller(true)       // 日志记录文件命名
	Logger.SetLevel(logrus.DebugLevel) // 日志输出级别
	Logger.Out = ioutil.Discard        // 禁止 logrus 的输出

	// 配置文件输入 hook
	logFolderPath, err := filepath.Abs(os.Getenv("LOG_PATH"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	logPath := path.Join(logFolderPath, "default.log")

	fmt.Println(logPath)

	writer, err := rotatelogs.New(
		logPath+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(logPath),          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(14*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		log.Fatalf("config log file system logger error. %v", errors.WithStack(err))
	}

	// 不同级别输出
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.AddHook(lfHook)

	Log = Logger.WithFields(logrus.Fields{"type": "DEFAULT"})
	LogAccess = Logger.WithFields(logrus.Fields{"type": "ACCESS"})
}
