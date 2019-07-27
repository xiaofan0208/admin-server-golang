package main 

import  (
    "github.com/gin-gonic/gin"
    "admin-server-golang/base"
    "admin-server-golang/routers"
    _ "admin-server-golang/base/database"
    // "fmt"
    "os"
    adminRouters "admin-server-golang/admin/routers"
)

func main(){
    r := gin.Default()
    routers.InitRouters(r)

    InitArgs()
 
    port := base.AppConfig.String("", "httpport")
    r.Run( ":" + port )  
}

func InitArgs(){
    if len(os.Args) <= 1 {
        return
    }

    switch os.Args[1] {
    case "-initAdmin":  // 初始化 管理员
        adminRouters.InitAdmin()
        break
    case "-initArbc":   // 初始化 菜单权限等
        adminRouters.InitMenus()
        break
    }
    os.Exit(1)
}