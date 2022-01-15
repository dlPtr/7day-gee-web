package gee

import "net/http"

func SetMetaUtf8() HandlerFunc {
	return func(c *Context) {
		c.Next()
		c.HTML(http.StatusOK, "<meta charset='utf-8'>")
	}
}
