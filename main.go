package main 

import  (
    "github.com/gin-gonic/gin"
	"admin-server-golang/base"
)


func main(){
	r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    port := base.AppConfig.String("", "httpport")
    r.Run( ":" + port )  
}