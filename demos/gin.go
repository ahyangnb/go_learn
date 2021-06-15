package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 设置模式
	gin.SetMode(gin.ReleaseMode)

	r.POST("/ping", func(c *gin.Context) {
		message := c.PostForm("message")
		fmt.Printf("message::%v\n", message)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
