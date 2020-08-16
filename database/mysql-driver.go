package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gogf/gf-demos/global"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct {
}

func (e *Mysql) Setup() {
	var err error
	dsn := g.Cfg().GetString("database.link")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)
	global.Eloquent, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		_ = gerror.Stack(err)
	}

	fmt.Println("数据库 连接成功")
	global.Eloquent.Logger.LogMode(logger.Info)
}
