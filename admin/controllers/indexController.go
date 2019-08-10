package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	// "net/http"
	// "fmt"
)

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index(c *gin.Context){
	data := gin.H{
		"Title" : "Dashboard",
	}
	ctl.DrawHTML( c , "admin/home/index.html" , data)
}  

// 注销
func (ctl *IndexController) Logout(c *gin.Context){
	session := sessions.Default(c)
	session.Set("BackendUser", nil)
	session.Save()

	ctl.Prepare(c)
}