package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/controllers"
)

func InitRouters(r *gin.Engine){
	r.GET("/index", (&controllers.IndexController{}).Index )
}