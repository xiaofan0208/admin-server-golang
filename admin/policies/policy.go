package policies

import (
	"github.com/gin-gonic/gin"
)

// 后台身份验证
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
 