package controllers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/models"
	// "encoding/json"
	// "fmt"
	"strconv"
)

// 菜单管理
type MenuAdminController struct {
	BaseController
}

func (ctl *MenuAdminController) Index(c *gin.Context){

	data := gin.H{
		"Title" : "菜单管理",
	}
	ctl.DrawHTML( c , "admin/home/menus.html" , data)

}

func (ctl *MenuAdminController) PostList(c *gin.Context) {
	query := make(map[string]interface{})
	filterMap := make(map[string]interface{})
	order := make(map[string]string)

	limit := c.DefaultPostForm("limit", "20") 
	limitInt64, _ := strconv.ParseInt(limit, 10, 64)

	offset := c.DefaultPostForm("offset", "0") 
	offsetInt64, _ := strconv.ParseInt(offset, 10, 64)

	menus , err := models.GetAllAdminMenusByQuery(query , filterMap ,  order , limitInt64 , offsetInt64)
	if err != nil {

	}
	ctl.jsonResult(c , models.JRCodeSucc  ,"" ,  menus)
}