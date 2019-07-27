package models


// 管理员 角色
type AdminRoles struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Name        string `form:"name" json:"name"`
	Pid         int    `form:"pid"  json:"pid"`
	Created     int64  `form:"created"  json:"created"`
	Status      uint8  `form:"status"   json:"status"`  
}

func(u AdminRoles) TableName() string {
	return "gin_admin.admin_roles"
}