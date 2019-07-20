package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (ctl *IndexController) Index(c *gin.Context){
	c.HTML(http.StatusOK, "admin/home/index.html", gin.H{
		"Title" : "Index Page",
	})
}