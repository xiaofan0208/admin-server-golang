package routers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/controllers"
	"admin-server-golang/admin/policies"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"admin-server-golang/base"
	"time"
	"admin-server-golang/admin/models"
	"admin-server-golang/admin/untils"
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

		v1.Use(policies.Authorize(router))
		
		v1.GET("/index", indexCtls.Index )    // 首页
		v1.GET("/logout", indexCtls.Logout )  // 注销
 
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

// 初始化管理员账号
func InitAdmin(){
	record := models.BackendUser{
		Username : "admin",
		Password : "123456",
		IsAdmin  : true,
	}
	user1 , _ := models.GetBackendUser(&record)
	if user1 != nil {
		return
	}
	_ ,err := models.CreateBackendUser(&record)
	if err != nil {

	}
}

// 初始化菜单
func InitMenus(){
	// 系统管理
	systemManage := &models.AdminMenus{ Name: "系统管理" , Type: models.MENU , Icon:"fa-adjust"  }
	// 角色管理
	roleManage :=  &models.AdminMenus{ Name: "角色管理" , Type: models.BUTTON , Icon:"fa-asterisk" , Url : "" }
	// 菜单管理
	menuManage :=  &models.AdminMenus{ Name: "菜单管理" , Type: models.BUTTON , Icon:"fa-bar-chart-o" , Url : "" }
	permissionManage :=  &models.AdminMenus{ Name: "权限管理" , Type: models.BUTTON , Icon:"fa-beer" , Url : "" }

	systemManage.Sons = []*models.AdminMenus{ roleManage , menuManage , permissionManage}


	// 用户管理
	usersManage := &models.AdminMenus{ Name: "用户管理" , Type: models.MENU , Icon:"fa-barcode"  }
	personlmenuManage :=  &models.AdminMenus{ Name: "修改个人资料" , Type: models.BUTTON , Icon:"fa-flask" , Url : "" }
	usersManage.Sons = []*models.AdminMenus{ personlmenuManage }

	untils.RegisterMenu(systemManage)
	untils.RegisterMenu(usersManage)	

	untils.InitMenu()
}