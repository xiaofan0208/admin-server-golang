package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/controllers"
	"admin-server-golang/admin/policies"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"admin-server-golang/base"
	"time"
)

func InitRouters(router *gin.Engine) {
	initSession(router)

	loginCtls := new(controllers.LoginController)
	indexCtls := new(controllers.IndexController)

    v1 := router.Group("/admin")
    {
		// 登录
		v1.GET("/login", loginCtls.Index )   
		v1.POST("/login", loginCtls.DoLogin ) 
		v1.POST("/register", loginCtls.Register ) 

		v1.Use(policies.Authorize())
		
		v1.GET("/index", indexCtls.Index )  
		
 
    }

}

func initSession(router *gin.Engine){
	// 初始化session
	secret := base.AppConfig.String("", "SESSION_KEY")
	store := cookie.NewStore([]byte( secret ))
	store.Options(sessions.Options{
        MaxAge: int(30 * time.Minute), //30min
        Path:   "/",
    })
	router.Use(sessions.Sessions("mysession", store))

}