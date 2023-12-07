package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	// 读取的文件名
	viper.SetConfigName("app")
	// 读取的文件类型
	viper.SetConfigType("yml")
	// 读取的路径
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

//func InitMySQL() {
//	//自定义模板打印SQL语句
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold: time.Second, //慢SQL阀值
//			LogLevel:      logger.Info, //级别
//			Colorful:      true,        //彩色
//		},
//	)
//
//	dns := viper.GetString("mysql.dns")
//	DB, _ = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
//	fmt.Println("MySQL inited ...")
//}
