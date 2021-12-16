package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, urlPattern string, handler HandlerFunc) {
	key := method + "-" + urlPattern
	engine.router[key] = handler
}

func (engine *Engine) GET(urlPattern string, handler HandlerFunc) {
	engine.addRoute("GET", urlPattern, handler)
}

func (engine *Engine) POST(urlPattern string, handler HandlerFunc) {
	engine.addRoute("POST", urlPattern, handler)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, req)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND")
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
