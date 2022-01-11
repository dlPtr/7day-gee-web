package main

import (
	"7day-go-demo/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1> Hey, there!</h1>")
		log.Printf("/ visited")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hey, %s --%s", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hey, %s --%s", c.Params["name"], c.Path)
	})

	r.GET("/my/hobbies/is/*hobbies", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"hobbies": c.Param("hobbies"),
		})
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

	if err := r.Run(":9999"); err != nil {
		fmt.Println("err is", err)
		return
	}

	fmt.Println("finish...")
}
