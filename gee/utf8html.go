package gee

func SetMetaUtf8() HandlerFunc {
	return func(c *Context) {
		c.Next()
		c.HTMLRaw("<meta charset='utf-8'>")
	}
}
