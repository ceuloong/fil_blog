package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	// 看清楚引入包，包的差异存在方法差异
	"gorm.io/gorm"
)

// DB 用大写声明(可以全局访问)
var DB *gorm.DB

// InitDB 创建数据库连接池
func InitDB() *gorm.DB {
	env := viper.GetString("env")
	datasource := viper.Sub("datasource")
	if env == "local" {
		datasource = viper.Sub("datasource-local")
	}
	host := datasource.GetString("host")
	port := datasource.GetString("port")
	database := datasource.GetString("database")
	username := datasource.GetString("username")
	password := datasource.GetString("password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database)

	//自定义模板打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阀值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	//fmt.Printf("args: %s\n", args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("fail err mysql", err.Error())
	}
	fmt.Println("MySQL inited ...")
	// gorm 自动创建表,需要放入model层中的模型，比如 User{}
	//db.Set("gorm:table_options", "charset=utf8mb4").
	//	AutoMigrate(&models.TronScan{})

	// 进行赋值 否则会空指针
	DB = db
	return db
}

// GetDB 获取DB的示例
func GetDB() *gorm.DB {
	return DB
}
