package models

// 管理员
type BackendUser struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	Username    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	RememberMe  bool   `form:"remember_me" json:"remember_me"`   // 是否记住 
	Created     int64  `json:"created"`
	Status      uint8  `json:"status"`                           // 状态: 1:正常 2:删除
}

func(u BackendUser) TableName() string {
	return "backend_user"
}