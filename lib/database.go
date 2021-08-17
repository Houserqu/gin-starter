package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic("connect database fail") // TODO: Logger
	}

	DB = db
}
