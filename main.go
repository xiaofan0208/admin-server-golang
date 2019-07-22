package main 

import  (
    "github.com/gin-gonic/gin"
    "admin-server-golang/base"
    "admin-server-golang/routers"
    _ "admin-server-golang/base/database"
)

func main(){
    r := gin.Default()
    routers.InitRouters(r)


    port := base.AppConfig.String("", "httpport")
    r.Run( ":" + port )  
}