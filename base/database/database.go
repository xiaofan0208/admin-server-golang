package database

import (
	_ "github.com/go-sql-driver/mysql"
	"admin-server-golang/base"
	"github.com/jinzhu/gorm"
    "admin-server-golang/helper"
    "fmt"
)

var DB *gorm.DB

func init(){
    db_type := base.AppConfig.String("", "db_type")
    db_config := base.AppConfig.String(db_type, "db_config")

    var err error
    DB, err = gorm.Open(db_type,db_config )
    // defer DB.Close()
    helper.CheckErr(err)

    fmt.Println("----database--")

    DB.DB().SetMaxIdleConns(10)   // 用于设置闲置的连接数
    DB.DB().SetMaxOpenConns(100)  // 用于设置最大打开的连接数

    DB.LogMode(true)  // 启用Logger，显示详细日志
}