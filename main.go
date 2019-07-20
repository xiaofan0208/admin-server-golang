package main 

import  (
    "github.com/gin-gonic/gin"
    "admin-server-golang/base"
    "admin-server-golang/routers"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
 //   _ "github.com/jinzhu/gorm/dialects/mysql" 
  
    "admin-server-golang/helper"
)

func initDB(){
    db_type := base.AppConfig.String("", "db_type")
    db_config := base.AppConfig.String(db_type, "db_config")

    db, err := gorm.Open(db_type,db_config )
    defer db.Close()
    helper.CheckErr(err)
}

func main(){
    initDB()

    r := gin.Default()
    routers.InitRouters(r)


    port := base.AppConfig.String("", "httpport")
    r.Run( ":" + port )  
}