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

	r.GET("/animals", func(c *gee.Context) {
		c.HTML(http.StatusOK, "ohh, animals! 11")
	})

	// r.Get("/animals", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "ohh, animals! 22 ")
	// })

	animalGroup := r.Group("/animals")
	{
		animalGroup.Get("/monkey", func(c *gee.Context) {
			c.SetHeader("x-content-type-options", "nosniff")
			c.SetHeader("foo", "bar")
			c.HTML(http.StatusOK, "<html><body><img src='' alt='monkey' /></html></body>")
		})
		monkeyGroup := animalGroup.Group("/monkey")
		{
			monkeyGroup.Get("/golden", func(c *gee.Context) {
				c.HTML(http.StatusOK, "<head><meta charset='utf-8'></head>金丝猴")
			})
			monkeyGroup.Get("/king", func(c *gee.Context) {
				c.HTML(http.StatusOK, "<meta charset='utf-8'>孙悟空")
			})
		}
		animalGroup.Get("/pm", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<meta charset='utf-8'> 社畜、工贼")
		})
		pm := animalGroup.Group("/pm")
		{
			pm.Get("/sxf", func(c *gee.Context) {
				c.HTML(http.StatusOK, "pua")
			})
		}
	}

	foodGroup := r.Group("/food")
	{
		foodGroup.Get("/sichuan", func(c *gee.Context) {
			c.HTML(http.StatusOK, "So fucking hot")
		})
	}

	if err := r.Run(":9999"); err != nil {
		fmt.Println("err is", err)
		return
	}

	fmt.Println("finish...")
}
