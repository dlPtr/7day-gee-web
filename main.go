package main

import (
	"fmt"
	"go-gee/gee"
	geeMid "go-gee/gee/middlewares"
	mid "go-gee/middlewares"
	"net/http"
	"text/template"
)

func main() {
	r := gee.New()

	r.SetFuncMap(template.FuncMap{
		"nameFmt": func(name string) string { return "[" + name + "]" },
	})
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello, gee!</h1>")
	})
	r.Use(geeMid.Logger(), geeMid.FakeLogger(), geeMid.Recovery())

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
	}
	animalGroup.Use(mid.SetMetaUtf8())

	// Day6：注册静态资源、渲染html模板
	r.Static("/assets", "./static/")
	r.Static("/css", "./static/css/")

	r.Get("/hello", func(c *gee.Context) {
		c.HTMLRender(http.StatusOK, "hello.html", map[string]interface{}{"name": c.Query("name"), "age": 24})
	})

	r.Get("/panic", func(c *gee.Context) {
		names := []string{"zhangsan", "lisi"}
		c.HTML(http.StatusOK, names[3])
	})

	if err := r.Run(":9999"); err != nil {
		fmt.Println("err is", err)
		return
	}

	fmt.Println("finish...")
}
