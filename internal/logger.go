package internal

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var defaultLogger *log.Logger

func init() {
	// 获取日志目录绝对路径
	logFolderPath, err := filepath.Abs(os.Getenv("LOG_PATH"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 判断目录是否存在，不存在则创建目录
	_, err = os.Stat(logFolderPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logFolderPath, os.ModePerm)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// 打开日志文件
	writer, err := os.OpenFile(filepath.Join(logFolderPath, "default.log"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defaultLogger = log.New(io.MultiWriter(writer), "", log.Lshortfile|log.LstdFlags)
}

func Log(v ...interface{}) {
	defaultLogger.Println(v)
}
