package models

import (
	db "admin-server-golang/base/database"
	"admin-server-golang/helper"
	"github.com/jinzhu/gorm"
)

//  系统菜单
type  AdminMenus struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Name        string `form:"name" json:"name"`
	Pid         int    `form:"pid"  json:"pid"`
	Url         string `form:"url"  json:"url"`           // 资源URL
	Type        int    `form:"type"  json:"type"`         // 类型  1: panel   2 :菜单  3 ：按钮
	Icon        string `form:"icon"  json:"icon"`         // 菜单小图标
	Sort        int    `form:"sort"  json:"sort"`         // '排序, 序号越大, 越靠前',
	Note        string `form:"note"  json:"note"`         // 备注
	Created     int64  `form:"created"  json:"created"`
	Status      uint8  `form:"status"   json:"status"`  

	Sons       []*AdminMenus  //`form:"-" json:"-"`
}

func(u AdminMenus) TableName() string {
	return "gin_admin.admin_menus"
}


// 创建
func CreateAdminMenus(record *AdminMenus) (*AdminMenus ,error ) {
	record.Created = helper.GetMillisecond()
	record.Status = Status_Normal
	if err := db.DB.Create(&record).Error ;  err != nil {
		helper.CheckErr(err)
		return nil,err
	}
	return record , nil
}


func GetAllAdminMenus()  ( []*AdminMenus , error ) {
	menus := []*AdminMenus{}
	if err := db.DB.Find(&menus).Error ; err != nil {
		if err != gorm.ErrRecordNotFound {
			helper.CheckErr(err)
		}	
		return nil, err
	}
	return menus , nil 
}


func DeleteAdminMenusByID(id int)  ( *AdminMenus , error ){
	menu := AdminMenus{}
	if err := db.DB.Where("Id = ?", id).Delete(&menu).Error ; err != nil  {
		if err != gorm.ErrRecordNotFound {
			helper.CheckErr(err)
		}	
		return nil, err
	}
	return &menu , nil 
}