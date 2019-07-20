package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"admin-server-golang/admin/models"
	// "encoding/json"
)

// 登录页面
type LoginController struct {
}

func (ctl *LoginController) Index(c *gin.Context){

	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{
		"SubTitle" : "Login Page",
		"formUrl" : "/admin/login",
		"listUrl" : "/admin/index",
	})
}

func (ctl *LoginController) Login (c *gin.Context) {
	var user  models.BackendUser
	if err := c.ShouldBind(&user); err != nil {
		data := &models.JsonResult{
			Code : models.JRCodeFailed ,
			Msg :  err.Error() ,
			Obj : nil,
		 }
		c.JSON(http.StatusBadRequest, data  )
		return
	}

	data := &models.JsonResult{
		Code : models.JRCodeSucc ,
		Msg :  "登录成功" ,
		Obj : nil,
	 }

	c.JSON(http.StatusOK, data )
}

