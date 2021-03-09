package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	r := gin.Default() // 返回默认的路由引擎

	r.GET("/hello", sayHello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	r.Run()
}
