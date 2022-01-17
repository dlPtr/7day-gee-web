package middlewares

import (
	"go-gee/gee"
)

func SetMetaUtf8() gee.HandlerFunc {
	return func(c *gee.Context) {
		c.Next()
		header := c.Writer.Header()
		if header["Content-Type"][0] == "text/html" {
			c.Raw("<meta charset='utf-8'>")
		}
	}
}
