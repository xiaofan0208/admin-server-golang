package controllers

import (
	"github.com/gin-gonic/gin"
	"admin-server-golang/base"
	// "net/http" 
)

type IndexController struct {
	base.Controller
}

func (ctl *IndexController) Index(c *gin.Context){
}