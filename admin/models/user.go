package models

import (
	db "admin-server-golang/base/database"
	"admin-server-golang/helper"
	// "fmt"
	// "encoding/json"
	"github.com/jinzhu/gorm"
)

// 管理员
type BackendUser struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Email       string `form:"email" json:"email"`
	Username    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	RememberMe  bool   `form:"remember_me" json:"remember_me" `   // 是否记住 
	Created     int64  `form:"created"  json:"created"`
	Status      uint8  `form:"status" json:"status"`                           // 状态: 1:正常 2:删除
}

func(u BackendUser) TableName() string {
	return "gin_admin.backend_user"
}

//查询管理员
func  GetBackendUser(record *BackendUser) ( *BackendUser , error){
	user := BackendUser{}
	if err := db.DB.Where("username = ?" ,record.Username).First(&user).Error ;  err != nil{
		if err != gorm.ErrRecordNotFound {
			helper.CheckErr(err)
		}	
		return nil,err
	}
	return &user , nil
}

func  GetBackendUserById(id int) (*BackendUser ,error ){
	user := BackendUser{}
	if err := db.DB.First(&user, id).Error ; err != nil {
		if err != gorm.ErrRecordNotFound {
			helper.CheckErr(err)
		}	
		return nil, err
	}
	return &user , nil
}

func CheckBackendUserByName(username , password string) (*BackendUser ,error ){
	user := BackendUser{}
	if err := db.DB.Where("username = ? and password = ? " ,username , password ).First(&user).Error ;  err != nil {
		if err != gorm.ErrRecordNotFound {
			helper.CheckErr(err)
		}	
		return nil, err
	}
	return &user , nil
}

// 创建
func CreateBackendUser(record *BackendUser) (*BackendUser ,error ) {
	record.Created = helper.GetMillisecond()
	record.Status = Status_Normal
	if err := db.DB.Create(&record).Error ;  err != nil {
		helper.CheckErr(err)
		return nil,err
	}
	return record , nil
}