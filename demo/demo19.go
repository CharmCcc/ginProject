package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义中间件
func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test4", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)

	})
	r.Run(":8081")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")

		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}
