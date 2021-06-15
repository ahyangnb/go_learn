package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login Binding as Json
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type ErrorBody struct {
	Code int
	Msg  string
}

func main() {
	r := gin.Default()

	// Set the Mode
	gin.SetMode(gin.ReleaseMode)

	v1 := r.Group("/v1")

	v1.POST("/ping", func(c *gin.Context) {
		message := c.PostForm("message")
		fmt.Printf("message::%v\n", message)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1.POST("/loginForm", func(c *gin.Context) {

		user := c.PostForm("user")
		if user == "" {
			e := ErrorBody{
				Msg:  "Please input user",
				Code: 1,
			}
			c.JSON(http.StatusOK, e)
			return
		}

		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
