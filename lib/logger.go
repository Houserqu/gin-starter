package lib

import (
	"fmt"
	"io"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	writer, err := os.OpenFile("./default.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create log file failed: %v", err)
	}

	fmt.Println("init log")

	logger = log.New(io.MultiWriter(writer), "", log.Lshortfile|log.LstdFlags)
}

func Log(v ...interface{}) {
	logger.Println(v)
}
