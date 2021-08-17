package lib

import (
	"fmt"
	"io"
	"log"
	"os"
)

var defaultLogger *log.Logger

func init() {
	writer, err := os.OpenFile("./default.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create log file failed: %v", err)
	}

	fmt.Println("init log")

	defaultLogger = log.New(io.MultiWriter(writer), "", log.Lshortfile|log.LstdFlags)
}

func Log(v ...interface{}) {
	defaultLogger.Println(v)
}
