package untils

import (
	"admin-server-golang/admin/models"
	// "encoding/json"
	// "fmt"
	//"admin-server-golang/base"
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


// 获取 tree 结构的菜单
func GetAdminTreeGrid()  ([]*models.AdminMenus) {
	menus , _ := models.GetAllAdminMenus()
	treegrid := groupList2TreeGrid(menus)
	return treegrid
}

func groupList2TreeGrid(list []*models.AdminMenus)  ([]*models.AdminMenus) {
	result := make([]*models.AdminMenus ,0 )
	for _ , item := range list  {
		if item.Pid == 0 {
			children := resourceAddSons(item, list, item.Sons )
			item.Sons = children
            result = append(result , item)
		}
	}
	return result
}

func resourceAddSons(cur *models.AdminMenus , list , result []*models.AdminMenus) ([]*models.AdminMenus) {
    for _ , item := range list {
        if item.Pid == cur.Id {
            children := resourceAddSons(item, list, item.Sons )
            item.Sons = children
            result = append(result, item)
        }
    }
    return result
}


func  Resource2Html( treegrid  []*models.AdminMenus  ,   html  string )  (string){

	for _ , item := range treegrid {
		url := "#"
		class := "dropdown-toggle"
		if item.Type == models.BUTTON {   // 可点击按钮
			url = item.Url
			class = ""
		}

		html += `<li class="">
				<a href='` + url + `'  class="` + class + `">
					<i class="` + item.Icon  + `"></i>
					<span class="menu-text">`  +item.Name  + ` </span>
				`
		// 如果有子菜单则添加箭头
		if  item.Sons != nil  {
			html += `<b class="arrow fa fa-angle-down"></b>`
		}
		html += "</a>"

		if item.Sons != nil {
			html += `<ul class="submenu nav-hide" style="display: block;">`
			html = Resource2Html(item.Sons , html)
			html += `</ul>`
		}

		html += `</li>`
	}
	return html
}