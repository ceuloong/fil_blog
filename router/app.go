package router

import (
	"blog/apis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/index", apis.GetIndex)
	//r.GET("/user/getUserList", service.GetUserList)
	//r.GET("/user/createUser", service.CreateUser)

	return r
}
