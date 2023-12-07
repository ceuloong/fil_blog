package main

import (
	"blog/apis"
	"blog/common"
	"blog/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	config.InitConfig()
	common.InitDB()

	//timeToCreatDb()
	//start()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "localhost", "192.168.83.109"})

	r.GET("/index-correct", apis.GetIndexCorrect)

	r.GET("/index", apis.GetIndex)

	r.GET("/update-addresses", apis.UpdateAddresses)

	r.GET("/update-balance", apis.UpdateAddressesBalance)

	r.GET("/tron", apis.TronAddress)

	r.Run(":3000")

	//httputils.UpdateAddressesBalance(1700640108, "")
}

func main1() {
	// fmt.Println("公众号：Golang来啦")
	fmt.Println(time.Now())
	ticker1 := time.NewTicker(4 * time.Hour)
	for {
		curTime := <-ticker1.C
		fmt.Println(curTime)
		//start()
	}
}

//func start() {
//	httputils.Start()
//	//httputils.SaveNodes("f01900855")
//
//	// 保存控制地址
//	httputils.UpdateAddresses("")
//
//	httputils.UpdateNodes("f01900855")
//
//	// 更新控制地址余额
//	httputils.UpdateAddressesBalance()
//}

// 定时创建数据库
func timeToCreatDb() {
	for {
		now := time.Now()                                                                              //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 1)                                                                 //通过now偏移1小时
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location()) //获取下一个小时的日期
		t := time.NewTimer(next.Sub(now))
		fmt.Printf("距下次执行时间还有：%v\n", next.Sub(now)) //计算当前时间到下一小时的时间间隔，设置一个定时器
		<-t.C
		fmt.Printf("执行时间为：%v\n", time.Now())
		//以下为定时执行的操作
		//start()
	}
}
