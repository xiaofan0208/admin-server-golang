package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/controllers"
	adminrouters "admin-server-golang/admin/routers"
)

func InitRouters(router *gin.Engine){
	router.Static("/static", "./static") // 静态文件路径
	router.LoadHTMLGlob("views/**/**/*")   //渲染html页面


	// 初始化后台路由
	adminrouters.InitRouters(router)

	router.GET("/index", (&controllers.IndexController{}).Index )
}

