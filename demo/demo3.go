package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取 GET 参数

func main() {
	router := gin.Default()

	//  匹配格式为 /welcome?firstname=**&lastname=**
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		//  下面是c.Request.URL.Query().GET("lastname")的简写
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	router.Run()
}
