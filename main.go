package main

import (
	"7day-go-demo/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello, gee!</h1>")
	})
	r.Use(gee.Logger(), gee.FakeLogger())

	r.GET("/animals", func(c *gee.Context) {
		c.HTML(http.StatusOK, "ohh, animals! 11")
	})
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
				c.HTML(http.StatusOK, "金丝猴")
			})
			monkeyGroup.Get("/king", func(c *gee.Context) {
				c.HTML(http.StatusOK, "孙悟空")
			})
		}
		animalGroup.Get("/pm", func(c *gee.Context) {
			c.HTML(http.StatusBadRequest, "社畜、工贼")
		})
		pm := animalGroup.Group("/pm")
		{
			pm.Get("/sxf", func(c *gee.Context) {
				c.HTML(http.StatusOK, "pua")
			})
		}
		pm.Use(gee.SetMetaUtf8())
	}

	if err := r.Run(":9999"); err != nil {
		fmt.Println("err is", err)
		return
	}

	fmt.Println("finish...")
}
