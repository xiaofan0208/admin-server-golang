package untils

import (
	"admin-server-golang/admin/models"
	// "encoding/json"
	// "fmt"
)

var (
	Menus []*models.AdminMenus
)

func RegisterMenu(menu  *models.AdminMenus){
	Menus = append(Menus ,menu )
}

func InitMenu(){
	// 删除所有旧数据
	m , _ := models.GetAllAdminMenus()
	for _ , item := range m {
		models.DeleteAdminMenusByID(item.Id)
	}

	for _ , item := range Menus {
		Id := insert(item)
		if item.Sons != nil {
			child(Id , item.Sons )
		}
	}
}

// 返回Id
func insert(record *models.AdminMenus) (int){
	menu , err := models.CreateAdminMenus(record)
	if err != nil {
		return 0
	}
	return menu.Id
}

func child(pid int , data  []*models.AdminMenus ) {
	for _ , item := range data {
		item.Pid = pid
		Id := insert(item)
		if item.Sons != nil {
			child(Id , item.Sons)
		}
	}
}