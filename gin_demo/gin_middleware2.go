package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计时间
	since := time.Since(start)
	fmt.Println("程序用时： ", since)
}

func main() {
	r := gin.Default()
	r.Use(MyTime)
	shoppingGroup := r.Group("shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
