package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由分组
func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run()
}

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}
