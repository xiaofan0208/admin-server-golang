package controllers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/admin/models"
	"github.com/gin-contrib/sessions"
	"net/http"
	// "encoding/json"
	// "fmt"
)

type BaseController struct{
}


func (c *BaseController) jsonResult( cc *gin.Context, code models.JsonResultCode , msg string, obj interface{}){
	data := &models.JsonResult{code, msg, obj}
	cc.JSON(http.StatusOK ,  data  )
	return
}

func(c *BaseController)  setBackendUser2Session(cc *gin.Context, userId int) error {
	user , err := models.GetBackendUserById(userId)
	if err != nil {
		return err
	}
	// 登录成功保存session
	session := sessions.Default(cc)
	session.Set("BackendUser", user)
	session.Save()
	return nil
}