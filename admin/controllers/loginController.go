package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录页面
type LoginController struct {
}

func (ctl *LoginController) Index(c *gin.Context){
	
	c.HTML(http.StatusOK, "admin/home/login.html", gin.H{
		"Title" : "Login Page",
	})
}