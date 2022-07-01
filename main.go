package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	router := gin.Default()

	// 2.绑定路由规则，执行的函数。gin.Context，封装了request和response
	router.GET("/hello", func(c *gin.Context) { // 两个HTTP端点，/hello和/os
		c.JSON(200, gin.H{ // 返回JSON格式的消息
			"message": "Hello World!",
		})
	})
	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS) // 返回纯文本格式的当前操作系统名称
	})
	// 3.监听端口，默认8080
	router.Run(":5000")
}
