package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	router.GET("/test2", func(c *gin.Context) {
		fmt.Printf("****************" + c.Request.URL.Path + "*****")
		c.Request.URL.Path = "/test3"
		router.HandleContext(c)
	})

	router.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	router.Run()
}
