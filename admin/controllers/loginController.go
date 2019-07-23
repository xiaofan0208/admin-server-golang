package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"admin-server-golang/admin/models"
	// "encoding/json"
	// "fmt"
	"admin-server-golang/helper"
	"strings"
)

// 登录页面
type LoginController struct {
	BaseController
}

func (ctl *LoginController) Index(c *gin.Context){

	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{
		"SubTitle" : "Login Page",
		"formUrl" : "/admin/login",
		"listUrl" : "/admin/index",
	})
}

func (ctl *LoginController) DoLogin (c *gin.Context) {
	var record  = new(models.BackendUser)
	if err := c.ShouldBind(record); err != nil {
		ctl.jsonResult(c , models.JRCodeFailed ,err.Error() ,  nil)
		return
	}
	username := strings.TrimSpace(record.Username)
	password := strings.TrimSpace(record.Password)
	if len(username) == 0|| len(password) == 0 {
		ctl.jsonResult(c , models.JRCodeFailed  ,"用户名和密码不正确" ,  "")
		return
	}
	password = helper.String2md5(password)

	user , err  := models.CheckBackendUserByName(username , password)
	if user != nil && err == nil {  // 用户存在
		ctl.setBackendUser2Session(c , user.Id)
		ctl.jsonResult(c , models.JRCodeSucc  ,"登录成功" ,  "")
		return
	}else{
		ctl.jsonResult( c , models.JRCodeFailed  , "用户名或者密码错误", "")
		return
	}
}


// 注册用户
func (ctl *LoginController) Register (c *gin.Context) {
	email := strings.TrimSpace( c.PostForm("email")  )
	username := strings.TrimSpace( c.PostForm("username") )
	password := strings.TrimSpace( c.PostForm("password") )
	repassword := strings.TrimSpace( c.PostForm("repassword") )

	if len(email) == 0 {
		ctl.jsonResult(c , models.JRCodeFailed  ,"邮箱不正确" ,  "")
		return	
	}
	if len(username) == 0|| len(password) == 0 ||  len(repassword) == 0 {
		ctl.jsonResult(c , models.JRCodeFailed  ,"用户名和密码不正确" ,  "")
		return
	}
	if password != repassword {
		ctl.jsonResult(c , models.JRCodeFailed  ,"两次密码不一致" ,  "")
		return
	}

	password = helper.String2md5(password)
	record := models.BackendUser{
		Email : email,
		Username : username,
		Password : password,
	}

	user1 , _ := models.GetBackendUser(&record)
	if user1 != nil {
		ctl.jsonResult(c , models.JRCodeFailed  ,"用户名已存在" ,  "")
		return
	}

	_ ,err := models.CreateBackendUser(&record)
	if err != nil {
		ctl.jsonResult(c , models.JRCodeFailed  ,"注册失败" ,  "")
		return
	}
	ctl.jsonResult(c , models.JRCodeSucc  ,"登录成功" ,  "")
	return
}