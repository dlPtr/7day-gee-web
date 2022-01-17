package middlewares

import (
	"go-gee/gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func FakeLogger() gee.HandlerFunc {
	return func(c *gee.Context) {
		log.Printf("FUCK")
	}
}
