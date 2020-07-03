package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取路径中的参数
func main() {
	router := gin.Default()

	//    /user/:name   能够匹配/user/***， 除了/user/  or  /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//    /user/:name/*action  能够匹配/user/**/  也能匹配 /user/***/**
	//    如果没有路由器匹配到上面的规则/user/*** , 它会重定向到 /user/**/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
