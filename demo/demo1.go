package main

import "github.com/gin-gonic/gin"

// 使用 GET, POST, PUT, PATCH, DELETE, OPTIONS
func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	router.Run()
}
