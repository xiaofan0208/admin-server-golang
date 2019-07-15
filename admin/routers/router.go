package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/controllers"
	"admin-server-golang/admin/policies"
)

func InitRouters(router *gin.Engine) {

	loginCtls := new(controllers.LoginController)

    v1 := router.Group("/admin")
    {
		v1.Use(policies.Authorize())

        v1.GET("/login", loginCtls.Index )  // 登录
    }

}