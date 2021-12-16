package main

import (
	"7day-go-demo/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "custom: URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "custom: Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}