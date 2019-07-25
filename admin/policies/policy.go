package policies

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	// "admin-server-golang/admin/models"
)

// 后台身份验证
func Authorize(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		backendUser := session.Get("BackendUser") 

		if backendUser == nil {
			c.Request.URL.Path = "/admin/login"
			router.HandleContext(c)
		}

		c.Next()
	}
}
 