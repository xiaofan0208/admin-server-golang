package controllers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/models"
	"github.com/gin-contrib/sessions"
	"net/http"
	// "encoding/json"
	"fmt"
)

type BaseController struct{
	backendUser *models.BackendUser  // 已经登录的用户
}

func(ctl *BaseController) Prepare(cc *gin.Context){
	session := sessions.Default(cc)
	backendUser := session.Get("BackendUser") 

	if backendUser != nil {
		ctl.backendUser = backendUser.( *models.BackendUser  ) 
	}else{
		// 跳转 登录页面 
		cc.Redirect(http.StatusFound,"/admin/login")
	}
}

//   渲染html
func(ctl *BaseController) DrawHTML(cc *gin.Context , html string , data gin.H){
	ctl.Prepare(cc)

	data["BackendUser"] = ctl.backendUser

	cc.HTML(http.StatusOK, html, data )
}


func (c *BaseController) jsonResult( cc *gin.Context, code models.JsonResultCode , msg string, obj interface{}){
	data := &models.JsonResult{code, msg, obj}
	fmt.Println("----------jsonResult--------data = " , data)
	cc.JSON(http.StatusOK ,  data  )
	return
}

// 设置session
func(c *BaseController)  setBackendUser2Session(cc *gin.Context, userId int) error {
	user , err := models.GetBackendUserById(userId)
	if err != nil {
		return err
	}
	c.backendUser = user
	// 登录成功保存session
	session := sessions.Default(cc)
	session.Set("BackendUser", user)
	session.Save()
	return nil
}