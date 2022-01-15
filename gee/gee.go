package gee

import (
	"log"
	"net/http"
	"strings"
)

type (
	// Group implement
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	// Engine implement the interface of ServeHTTP
	Engine struct {
		router *router
		*RouterGroup
		groups []*RouterGroup // all groups
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	for _, eachGroup := range engine.groups {
		if strings.HasPrefix(req.URL.Path, eachGroup.prefix) {
			c.handlers = append(c.handlers, eachGroup.middlewares...)
		}
	}

	engine.router.handle(c)
}

// function implements of RouterGroup
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: g.engine,
	}

	g.engine.groups = append(g.engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) Get(pattern string, handler HandlerFunc) {
	// FIXME: "GET" is a magic string, so is "POST".
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) Post(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}
