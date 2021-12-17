package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1> Hey, there!</h1>")
		log.Printf("/ visited")
	})
	r.POST("/login", func(c *gee.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		code, msg, data := 1, "login failed, check your name and passwd", struct{}{}
		if username == "ysb" && password == "123456" {
			code, msg = 0, "login successfully"
		}

		c.JSON(http.StatusOK, gee.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
	})

	r.Run(":9999")
}
