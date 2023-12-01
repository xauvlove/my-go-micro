package main

import (
	"my-micro/web/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化路由
	router := gin.Default()

	// 路由匹配
	router.Static("/home", "view")

	rg1 := router.Group("/api/v1.0")
	{
		rg1.GET("/session", controller.GetSession)
		rg1.GET("/imagecode/:uuid", controller.GetImageCd)
		rg1.GET("/smscode/:phone", controller.GetSmsCd)
		rg1.POST("/users", controller.PostRet)
	}

	// 启动
	router.Run(":8080")
}
