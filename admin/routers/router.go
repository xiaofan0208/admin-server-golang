package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/controllers"
	"admin-server-golang/admin/policies"
)

func InitRouters(router *gin.Engine) {

	loginCtls := new(controllers.LoginController)
	indexCtls := new(controllers.IndexController)

    v1 := router.Group("/admin")
    {
		// 登录
		v1.GET("/login", loginCtls.Index )   
		v1.POST("/login", loginCtls.Login ) 

		v1.Use(policies.Authorize())
		
		v1.GET("/index", indexCtls.Index )  
		
 
    }

}